package geneticalgorithm

import (
	"math"
	"fmt"
)

type IsOptimized func(individuum Individuum) bool

type GeneticAlgorithm struct {
	Popsize       int
	MaxIterations uint64
	Mutator       Mutate
	Recombiner    Recombine
	Selector      Select
}

func (ga GeneticAlgorithm) Optimize(optimized IsOptimized, verbose bool) Individuum {

	pop := GenerateStartPopulation(ga.Popsize)

	i := uint64(0)
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
				return math.Exp(-individuum.GetPhenotype())
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

	channelSize := int(math.Ceil(float64(ga.Popsize) / 4))

	parents := make(chan Individuum, channelSize)
	children := make(chan Individuum, channelSize)
	mutated := make(chan Individuum, channelSize)
	selected := make(chan Individuum, channelSize)
	result := make(chan Individuum, 1)

	go pop.streamIndividuals(parents)
	go ga.Recombiner(parents, children)
	go ga.Mutator(children, mutated, MutateContext{0, 1})
	go ga.Selector(mutated, selected, int(ga.Popsize/2),
		func(individuum Individuum) float64 {
			return math.Exp(-individuum.GetPhenotype())
		})

	info := make(chan AnalyzeInfo, 10)
	go analyzeIndividuum(selected, parents, result, info, optimized, ga.MaxIterations)

	best := 1e9
	for f := range info {
		if newBest := f.BestPhenotype; verbose && newBest < best {
			best = newBest
			fmt.Printf("New best phenotype %.6f after %d individuals.\n", f.BestPhenotype, f.N)
		}
	}

	return <-result
}

type AnalyzeInfo struct {
	N                uint64
	BestPhenotype    float64
	CurrentPhenotype float64
}

func (ai AnalyzeInfo) String() string {
	return fmt.Sprintf("Individuum counter: %d, Best Phenotype: %.3f, Current Phenotype: %.3f", ai.N, ai.BestPhenotype, ai.CurrentPhenotype)
}

func analyzeIndividuum(in <-chan Individuum, out, result chan<- Individuum, info chan<- AnalyzeInfo, optimized IsOptimized, maxIterations uint64) {
	defer close(out)
	defer close(info)
	bestPhenotype := 1e9
	var bestIndividuum Individuum
	counter := uint64(0)
	for individuum := range in {
		out <- individuum
		phenotype := individuum.GetPhenotype();
		if phenotype < bestPhenotype {
			bestPhenotype = phenotype
			bestIndividuum = individuum
			if optimized(individuum) {
				result <- bestIndividuum
				info <- AnalyzeInfo{counter, bestPhenotype, phenotype}
				return
			}
		}
		info <- AnalyzeInfo{
			N:                counter,
			BestPhenotype:    bestPhenotype,
			CurrentPhenotype: phenotype,
		}
		if (counter > maxIterations) {
			result <- bestIndividuum
			return
		}
		counter++
	}
}
