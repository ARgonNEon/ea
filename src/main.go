package main

import (
	"fmt"
	"ga"
	"hillclimber"
	"ackley"
)

func main() {
	demonstrateGeneticAlgorithm()
}

func demonstrateHillClimber() {
	climber := hillclimber.Hillclimber{
		Nmax: 1000,
		Delta: 1e-1,
		Dim: 6,
	}
	result := climber.Climb([]float64{5, 1, 9, 10, 1.5, 2.33})
	fmt.Println("Hillclimber:")
	fmt.Println("\tResult: ", result)
	fmt.Println("\tValue: ", ackley.Ackley(result))
}

func demonstrateGeneticAlgorithm() {
	fmt.Println("Genetic Algorithm")
	g := ga.GeneticAlgorithm{
		Popsize: 50,
		MaxIterations: 20000,
		Mutator: ga.AdaptiveGaussianMutator,
		Recombiner: ga.OnePointCrossOver,
		Selector: ga.RemainderStochasticSampling,
	}
	result := g.Optimize(func (individuum ga.Individuum) bool {
		return individuum.GetFitness() < 0.05
	}, true)
	fmt.Println()
	fmt.Println("Result: ")
	fmt.Println(result)


}
