package evolutionstrategy

import "wesx.de/ArneS/ea/optimizer"

type EvolutionStrategy struct {
	Popsize    int
	Lambda     int
	UpperBound float64
	LowerBound float64
}

func (es EvolutionStrategy) Optimize(optimized optimizer.IsOptimized, verbose bool) optimizer.Individuum {
	return optimizer.MakeRandomIndividuum(6, -40, 40)
}
