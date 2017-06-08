package main

import (
	"bufio"
	"fmt"
	"os"

	"wesx.de/ArneS/ea/ackley"
	"wesx.de/ArneS/ea/evolutionstrategy"
	"wesx.de/ArneS/ea/geneticalgorithm"
	"wesx.de/ArneS/ea/hillclimber"
	"wesx.de/ArneS/ea/optimizer"
)

func main() {
	demonstrateHillClimber()
	waitKey()
	fmt.Println("-------------------------------------------------")
	demonstrateGeneticAlgorithm()
	waitKey()
	fmt.Println("-------------------------------------------------")
	demonstratePipelinedGeneticAlgorithm()
	waitKey()
	fmt.Println("-------------------------------------------------")
	demonstrateEvolutionStategy()
}

func waitKey() {
	fmt.Println("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func demonstrateHillClimber() {
	climber := hillclimber.Hillclimber{
		Nmax:  1000,
		Delta: 1e-1,
		Dim:   6,
	}
	result := climber.Climb([]float64{5, 1, 9, 10, 1.5, 2.33})
	fmt.Println("Hillclimber:")
	fmt.Println("\tResult: ", result)
	fmt.Println("\tValue: ", ackley.Ackley(result))
	fmt.Println()
}

func demonstrateGeneticAlgorithm() {
	fmt.Println("Genetic Algorithm")
	g := geneticalgorithm.GeneticAlgorithm{
		Popsize:       50,
		MaxIterations: 25000,
		Mutator:       geneticalgorithm.NonUniformMutator,
		Recombiner:    geneticalgorithm.OnePointCrossOver,
		Selector:      geneticalgorithm.RemainderStochasticSampling,
	}
	result := g.Optimize(func(individuum optimizer.Fitnessable) bool {
		return individuum.GetPhenotype() < 0.05
	}, true)
	fmt.Println()
	fmt.Println("Result: ")
	fmt.Println(result)
	fmt.Println()
}

func demonstrateEvolutionStategy() {
	fmt.Println("Evolution Strategy")
	e := evolutionstrategy.EvolutionStrategy{
		Popsize:       100,
		Lambda:        7,
		MaxIterations: 1000,
		UpperBound:    40.0,
		LowerBound:    -40.0,
	}
	result := e.Optimize(func(individuum optimizer.Fitnessable) bool {
		return individuum.GetPhenotype() < 1e-9
	}, true)
	fmt.Println()
	fmt.Println("Result: ")
	fmt.Println(result)
	fmt.Println()

}

func demonstratePipelinedGeneticAlgorithm() {
	fmt.Println("Genetic Algorithm (Pipelined)")
	g := geneticalgorithm.GeneticAlgorithm{
		Popsize:       100,
		MaxIterations: 1e8,
		Mutator:       geneticalgorithm.NonUniformMutator,
		Recombiner:    geneticalgorithm.OnePointCrossOver,
		Selector:      geneticalgorithm.RemainderStochasticSampling,
	}
	result := g.OptimizePipelined(func(individuum optimizer.Fitnessable) bool {
		return individuum.GetPhenotype() < 0.05
	}, true)
	fmt.Println()
	fmt.Println("Result: ")
	fmt.Println(result)
	fmt.Println()
}
