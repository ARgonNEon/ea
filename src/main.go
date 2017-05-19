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
	climber := hillclimber.Hillclimber{1000, 1e-1, 6}
	result := climber.Climb([]float64{5, 1, 9, 10, 1.5, 2.33})
	fmt.Println("Hillclimber:")
	fmt.Println("\tResult: ", result)
	fmt.Println("\tValue: ", ackley.Ackley(result))
}

func demonstrateGeneticAlgorithm() {
	fmt.Println("Genetic Algorithm")
	pop := ga.GenerateStartPopulation(50)
	fmt.Printf("")
	pop.Analyze()
}