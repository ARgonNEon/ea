package evolutionstrategy

import (
	"fmt"

	"wesx.de/ArneS/ea/optimizer"
)

type EvolutionStrategy struct {
	Popsize       int
	Lambda        int
	MaxIterations int
	UpperBound    float64
	LowerBound    float64
}

func (es EvolutionStrategy) Optimize(optimized optimizer.IsOptimized, verbose bool) optimizer.Individuum {
	pop := MakePopulation(es.Popsize)
	startPop := make(chan optimizer.Individuum)
	go es.createStartIndividuals(startPop)
	pop.collectIndividuals(startPop)
	fmt.Println(pop)

	mutator := Mutator{
		Lambda:     es.Lambda,
		LowerBound: es.LowerBound,
		UpperBound: es.UpperBound,
		Sigma:      1,
	}

	selector := Selector{
		Popsize: es.Popsize,
	}

	for i := 0; i < es.MaxIterations; i++ {
		initial := make(chan optimizer.Individuum)
		mutated := make(chan optimizer.Individuum)
		selected := make(chan optimizer.Individuum)

		go pop.streamIndividuals(initial)
		go mutator.Identity(initial, mutated)
		go selector.Select(mutated, selected)
		pop.collectIndividuals(selected)
		if verbose {
			fmt.Println(pop)
		}
	}
	return pop.findBest()
}

func (es EvolutionStrategy) createStartIndividuals(out chan<- optimizer.Individuum) {
	for i := 0; i < es.Popsize; i++ {
		out <- optimizer.MakeRandomIndividuum(6, es.LowerBound, es.UpperBound)
	}
}
