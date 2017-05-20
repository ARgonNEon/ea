package ga

import "math"

//import "math"

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

	channelSize := int(math.Min(100, float64(ga.Popsize)))
	i := 0
	for {

		parents := make(chan Individuum, channelSize)
		children := make(chan Individuum, channelSize)
		mutated := make(chan Individuum, channelSize)
		selected := make(chan Individuum, channelSize)

		go pop.streamIndividuals(parents)
		go ga.Recombiner(parents, children)
		go ga.Mutator(children, mutated, MutateContext{pop.age, ga.MaxIterations})
		go ga.Selector(mutated, selected, ga.Popsize,
			func(individuum Individuum) float64 {
				return math.Exp(-individuum.GetFitness())
			})
		pop.collectIndividuals(selected)
		best, _ := pop.findBest()
		pop.Analyze()
		if i++; optimized(best) || i>=ga.MaxIterations{
			return best
		}
	}
}
