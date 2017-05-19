package ga

import "math"

type Optimizer interface {
	Optimize() Individuum
}

type GeneticAlgorithm struct {
	population Population
	mutator    Mutator
	recombiner Recombiner
	selector   Selector
}

func MakeGeneticAlgorithm() GeneticAlgorithm {
	return GeneticAlgorithm{GenerateStartPopulation(Popsize), nil, nil, nil}
}

func (ga GeneticAlgorithm) Optimize() Individuum {

	pop := ga.population
	for i := 0; i < Iterations; i++ {
		pop = ga.recombiner.Recombine(pop)

		pop = ga.selector.Select(pop, func(individuum Individuum) float64 {
			return math.Exp(-individuum.getFitness())
		})
	}
	return nil
}
