package ga

import "math"

//import "math"

type Optimizer interface {
	Optimize() Individuum
}

type GeneticAlgorithm struct {
	mutator    Mutator
	recombiner Recombiner
	selector   Selector
}

func MakeGeneticAlgorithm() GeneticAlgorithm {
	return GeneticAlgorithm{nil, nil, nil}
}

func (ga GeneticAlgorithm) Optimize() Individuum {

	pop := GenerateStartPopulation(Popsize)

	parents := make(chan Individuum)
	children := make(chan Individuum)
	mutated := make(chan Individuum)
	selected := make(chan Individuum)
	quit := make(chan bool)


	target_function := func(individuum Individuum) float64 {
			return math.Exp(-individuum.getFitness())
	}

	for i := 0; i < Iterations; i++ {
		go pop.streamIndividuals(parents, quit)
		go ga.recombiner.Recombine(parents, children)
		go ga.mutator.Mutate(children, mutated)
		go ga.selector.Select(mutated, selected, target_function)
		pop.collectIndividuals(selected)
		quit <- true
	}
	return nil
}
