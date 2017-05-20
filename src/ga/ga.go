package ga

import "math"

//import "math"

type Optimizer interface {
	Optimize() Population
}

type GeneticAlgorithm struct {
	Popsize	int
	Iterations int
	mutator    Mutate
	recombiner Recombine
	selector   Select
}

func MakeGeneticAlgorithm(popsize, iterations int) GeneticAlgorithm {
	return GeneticAlgorithm{popsize,
		iterations,
		AdaptiveGaussianMutator,
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