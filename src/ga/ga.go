package ga

import "math"

//import "math"

type Optimizer interface {
	Optimize() Individuum
}

type GeneticAlgorithm struct {
	mutator    Mutate
	recombiner Recombine
	selector   Select
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

	for i := 0; i < Iterations; i++ {
		go pop.streamIndividuals(parents, quit)
		go ga.recombiner(parents, children)
		go ga.mutator(children, mutated)
		go ga.selector(mutated, selected,
			func(individuum Individuum) float64 {
				return math.Exp(-individuum.getFitness())
			})
		pop.collectIndividuals(selected)
		quit <- true
		pop.Analyze()
	}
	return nil
}
