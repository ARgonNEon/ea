package geneticalgorithm

import (
	"math"
	"fmt"
)

type IsOptimized func(individuum Individuum) bool

type GeneticAlgorithm struct {
	Popsize       int
	MaxIterations int
	Mutator       Mutate
	Recombiner    Recombine
	Selector      Select
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

func (ga GeneticAlgorithm) OptimizePipelined(optimized IsOptimized, verbose bool) Individuum {
	pop := GenerateStartPopulation(ga.Popsize)

	channelSize := ga.Popsize / 4

	ancient := make(chan Individuum)
	parents := make(chan Individuum, channelSize)
	children := make(chan Individuum, channelSize)
	mutated := make(chan Individuum, channelSize)
	selected := make(chan Individuum, channelSize)
	loopback := make(chan Individuum, channelSize)
	result := make(chan Individuum, 1)

	go pop.streamIndividuals(ancient)
	go loopbackTee(ancient, loopback, parents)
	go ga.Recombiner(parents, children)
	go ga.Mutator(children, mutated, MutateContext{0, 1})
	go ga.Selector(mutated, selected, int(ga.Popsize/2),
		func(individuum Individuum) float64 {
			return math.Exp(-individuum.GetFitness())
		})

	info := make(chan AnalyzeInfo, 10)
	go analyzeIndividuum(selected, loopback, result, info, optimized)

	best := 1e9
	for f := range info {
		if verbose && f.BestFitness < best {
			best = f.BestFitness
			fmt.Printf("New best fitness %.6f after %d individuals.\n", f.BestFitness, f.N)
		}
	}

	return <-result
}

type AnalyzeInfo struct {
	N              int64
	BestFitness    float64
	CurrentFitness float64
}

func (ai AnalyzeInfo) String() string {
	return fmt.Sprintf("Individuum counter: %d, Best Fitness: %.3f, Current Fitness: %.3f", ai.N, ai.BestFitness, ai.CurrentFitness)
}

func analyzeIndividuum(in <-chan Individuum, out, result chan<- Individuum, info chan<- AnalyzeInfo, optimized IsOptimized) {
	defer close(out)
	defer close(info)
	best := 1e9
	counter := int64(0)
	for individuum := range in {
		fitness := individuum.GetFitness();
		if fitness < best {
			best = fitness
			if optimized(individuum) {
				result <- individuum
				info <- AnalyzeInfo{counter,best,fitness}
				return
			}
		}
		info <- AnalyzeInfo{
			N:              counter,
			BestFitness:    best,
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
