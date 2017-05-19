package main

import "fmt"
import "ga"


func main() {
	pop := ga.GenerateStartPopulation(50)
	fmt.Printf("")
	pop.Analyze()
}
