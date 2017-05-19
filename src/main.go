package main

import "fmt"
import "ga"
import (
	"hillclimber"
	"ackley"
)

func main() {
	demonstrateHillClimber()
}

func demonstrateHillClimber() {
	climber := hillclimber.Hillclimber{1000, 1e-1, 6}
	result := climber.Climb([]float64{1, 1, 1, 0.5, 0.5, 0.5})
	fmt.Println("Result: ", result)
	fmt.Println("Value: ", ackley.Ackley(result))
}

func demonstrateGeneticAlgorithm() {
	pop := ga.GenerateStartPopulation(50)
	fmt.Printf("")
	pop.Analyze()
}