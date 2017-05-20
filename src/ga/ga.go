package ga

import (
	"math"
	"fmt"
)

//import "math"

type Optimizer interface {
	Optimize() Population
}

type GeneticAlgorithm struct {
	Popsize    int
	Iterations int
	mutator    Mutate
	recombiner Recombine
	selector   Select
}

func MakeGeneticAlgorithm(popsize, iterations int) GeneticAlgorithm {
	return GeneticAlgorithm{popsize,
				iterations,
				NonUniformMutator,
				OnePointCrossOver,
				RemainderStochasticSampling}
}

func (ga GeneticAlgorithm) Optimize() Population {

	pop := GenerateStartPopulation(ga.Popsize)

	channelSize := int(math.Min(100, float64(ga.Popsize)))

	for i := 1; i < ga.Iterations; i++ {

		parents := make(chan Individuum, channelSize)
		children := make(chan Individuum, channelSize)
		mutated := make(chan Individuum, channelSize)
		selected := make(chan Individuum, channelSize)

		go pop.streamIndividuals(parents)
		go ga.recombiner(parents, children)
		go ga.mutator(children, mutated, MutateContext{pop.age, ga.Iterations})
		go ga.selector(mutated, selected, ga.Popsize,
			func(individuum Individuum) float64 {
				return math.Exp(-individuum.getFitness())
			})
		pop.collectIndividuals(selected)
	}
	return pop
}

func (ga GeneticAlgorithm) OptimizePipelined() Individuum {
	pop := GenerateStartPopulation(ga.Popsize)

	ancient := make(chan Individuum, ga.Popsize)
	parents := make(chan Individuum, ga.Popsize)
	children := make(chan Individuum, ga.Popsize)
	mutated := make(chan Individuum, ga.Popsize)
	selected := make(chan Individuum, ga.Popsize)
	loopback := make(chan Individuum, ga.Popsize)
	result := make(chan Individuum)

	go pop.streamIndividuals(ancient)
	go loopbackTee(ancient, loopback, parents)
	go ga.recombiner(parents, children)
	go ga.mutator(children, mutated, MutateContext{0, 1})
	go ga.selector(mutated, selected, int(ga.Popsize/2),
		func(individuum Individuum) float64 {
			return math.Exp(-individuum.getFitness())
		})

	info := make(chan AnalyzeInfo, 10)
	go analyzeIndividuum(selected, loopback, result, info)

	for f := range info {
		fmt.Println(f)
	}

	return nil
}

type AnalyzeInfo struct {
	N           int64
	BestFitness float64
	CurrentFitness float64
}

func (ai AnalyzeInfo) String() string {
	return fmt.Sprintf("Individuum counter: %d, Best Fitness: %.3f, Current Fitness: %.3f", ai.N, ai.BestFitness, ai.CurrentFitness)
}

func analyzeIndividuum(in <-chan Individuum, out, result chan<- Individuum, info chan<- AnalyzeInfo) {
	best := 1e9
	counter := int64(0)
	for individuum := range in {
		fitness := individuum.getFitness();
		if fitness < best {
			best = fitness
		}
		info <- AnalyzeInfo{
			N:           counter,
			BestFitness: best,
			CurrentFitness: fitness,
		}
		out <- individuum
		counter++
	}
}

func loopbackTee(input, loopback <-chan Individuum, output chan<- Individuum) {
	defer close(output)
	for individuum := range input {
		output <- individuum
	}
	for individuum := range loopback {
		output <- individuum
	}
}
