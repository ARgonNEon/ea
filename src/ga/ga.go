package ga

import (
	"math"
	"fmt"
)

//import "math"

type Optimizer interface {
	Optimize() Population
}

type GeneticAlgorithm struct {
	Popsize	int
	Iterations int
	mutator    Mutate
	recombiner Recombine
	selector   Select
}

func MakeGeneticAlgorithm(popsize, iterations int) GeneticAlgorithm {
	return GeneticAlgorithm{popsize,
		iterations,
		AdaptiveGaussianMutator,
		OnePointCrossOver,
		RemainderStochasticSampling}
}


func (ga GeneticAlgorithm) Optimize() Population {

	pop := GenerateStartPopulation(ga.Popsize)

	channelSize := int(math.Min(100, float64(ga.Popsize)))

	for i := 1; i < ga.Iterations; i++ {

		parents := make(chan Individuum, channelSize)
		children := make(chan Individuum, channelSize)
		mutated := make(chan Individuum, channelSize)
		selected := make(chan Individuum, channelSize)
		quit := make(chan bool)

		go pop.streamIndividuals(parents, quit)
		go ga.recombiner(parents, children, ga.Popsize)
		go ga.mutator(children, mutated, MutateContext{pop.age, ga.Iterations})
		go ga.selector(mutated, selected, ga.Popsize,
			func(individuum Individuum) float64 {
				return math.Exp(-individuum.getFitness())
			})
		pop.collectIndividuals(selected)
		quit <- true
	}
	return pop
}

func (ga GeneticAlgorithm) OptimizePipelined() Individuum {
	pop := GenerateStartPopulation(ga.Popsize)

	channelSize := ga.Popsize

	ancient := make(chan Individuum, channelSize)
	quitAncient := make(chan bool)
	loopback := make(chan Individuum, channelSize)
	children := make(chan Individuum, channelSize)
	mutated := make(chan Individuum, channelSize)
	selected := make(chan Individuum, channelSize)
	analyzed := make(chan Individuum, channelSize)
	stop := make(chan bool)

	go pop.streamIndividuals(ancient, quitAncient)
	go streamAndLoopback(ancient, analyzed, loopback, stop, quitAncient, ga.Popsize)
	go ga.recombiner(loopback, children, 0)
	go ga.mutator(children, mutated, MutateContext{0, 1})
	go ga.selector(mutated, selected, 50,
		func(individuum Individuum) float64 {
			return math.Exp(-individuum.getFitness())
		})

	debug := make(chan float64)
	go analyzeIndividuum(selected, loopback, debug)

	for f := range debug {
		fmt.Println(f)
	}


	return nil
}

func streamAndLoopback(ancient, loopback <-chan Individuum, out chan<-Individuum, stop <-chan bool, quitAncient chan<- bool, nAncients int) {
	for i:=0; i<nAncients; i++ {
		out <- <-ancient
	}
	quitAncient <- true
	for {
		select {
		case next := <-loopback:
			out <- next
		case <- stop:
			return
		}
	}
}

func analyzeIndividuum(in <-chan Individuum, out chan<- Individuum, chBest chan<- float64) {
	best := 1e9
	for individuum := range in {

		if fitness := individuum.getFitness(); fitness < best {
			best = fitness
			chBest <- best
		}
		out <- individuum
	}
}