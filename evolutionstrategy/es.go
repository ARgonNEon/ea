package evolutionstrategy

import "wesx.de/ArneS/ea/optimizer"

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
	es.createStartIndividuals(startPop)
	pop.collectIndividuals(startPop)

	mutator := Mutator{
		Lambda:     es.Lambda,
		LowerBound: es.LowerBound,
		UpperBound: es.UpperBound,
		Sigma:      1,
	}

	initial := make(chan optimizer.Individuum)
	mutated := make(chan optimizer.Individuum)
	for i := 0; i < es.MaxIterations; i++ {
		go pop.streamIndividuals(initial)
		go mutator.Mutate(initial, mutated)

	}
}

func (es EvolutionStrategy) createStartIndividuals(out chan<- optimizer.Individuum) {
	for i := 0; i < es.Popsize; i++ {
		out <- optimizer.MakeRandomIndividuum(6, es.LowerBound, es.UpperBound)
	}
}
