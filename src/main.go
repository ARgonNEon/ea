package main

import (
	"fmt"

	"ga"
)

func main() {
	pop := ga.GenerateStartPopulation(50)
	fmt.Printf("")
	pop.Analyze()
}
