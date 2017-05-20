package main

import (
	"fmt"
	"ga"
	"hillclimber"
	"ackley"
)

func main() {
	demonstrateHillClimber()
	fmt.Println("-------------------------------------------------")
	demonstrateGeneticAlgorithm()
	fmt.Println("-------------------------------------------------")
	demonstratePipelinedGeneticAlgorithm()
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
	fmt.Println()
}

func demonstrateGeneticAlgorithm() {
	fmt.Println("Genetic Algorithm")
	g := ga.GeneticAlgorithm{
		Popsize: 50,
		MaxIterations: 25000,
		Mutator: ga.AdaptiveGaussianMutator,
		Recombiner: ga.OnePointCrossOver,
		Selector: ga.RemainderStochasticSampling,
	}
	result := g.Optimize(func (individuum ga.Individuum) bool {
		return individuum.GetFitness() < 0.05
	}, false)
	fmt.Println()
	fmt.Println("Result: ")
	fmt.Println(result)
	fmt.Println()
}

func demonstratePipelinedGeneticAlgorithm() {
	fmt.Println("Genetic Algorithm (Pipelined)")
	g := ga.GeneticAlgorithm{
		Popsize: 50,
		MaxIterations: 1e9,
		Mutator: ga.NonUniformMutator,
		Recombiner: ga.OnePointCrossOver,
		Selector: ga.RemainderStochasticSampling,
	}
	result := g.OptimizePipelined(func (individuum ga.Individuum) bool {
		return individuum.GetFitness() < 0.05
	}, false)
	fmt.Println()
	fmt.Println("Result: ")
	fmt.Println(result)
	fmt.Println()
}
