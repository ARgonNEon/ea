package ga

import (
	"math"
	"fmt"
)

type IsOptimized func(individuum Individuum) bool

type GeneticAlgorithm struct {
	Popsize	int
	MaxIterations int
	Mutator    Mutate
	Recombiner Recombine
	Selector   Select
}

func (ga GeneticAlgorithm) Optimize(optimized IsOptimized, verbose bool) Individuum {

	pop := GenerateStartPopulation(ga.Popsize)

	i := 0
	for {

		parents := make(chan Individuum)
		children := make(chan Individuum)
		mutated := make(chan Individuum)
		selected := make(chan Individuum)

		go pop.streamIndividuals(parents)
		go ga.Recombiner(parents, children)
		go ga.Mutator(children, mutated, MutateContext{pop.age, ga.MaxIterations})
		go ga.Selector(mutated, selected, ga.Popsize,
			func(individuum Individuum) float64 {
				return math.Exp(-individuum.GetFitness())
			})
		pop.collectIndividuals(selected)
		best, _ := pop.findBest()
		if verbose {
			fmt.Println(pop)
		}
		if i++; optimized(best) || i >= ga.MaxIterations {
			return best
		}
	}
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
