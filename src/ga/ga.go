package ga

import "math"

//import "math"

type Optimizer interface {
	Optimize() Individuum
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
		DummyMutator,
		DummyRecombiner,
		DummySelector}
}

func (ga GeneticAlgorithm) Optimize() Individuum {

	pop := GenerateStartPopulation(ga.Popsize)

	for i := 1; i < ga.Iterations; i++ {
		parents := make(chan Individuum)
		children := make(chan Individuum)
		mutated := make(chan Individuum)
		selected := make(chan Individuum)
		quit := make(chan bool)

		go pop.streamIndividuals(parents, quit)
		go ga.recombiner(parents, children, ga.Popsize)
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
