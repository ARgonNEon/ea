package ga

//import "math"

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

	//pop := ga.population

	/*target_function := func(individuum Individuum) float64 {
			return math.Exp(-individuum.getFitness())
	}*/

	for i := 0; i < Iterations; i++ {

	}
	return nil
}
