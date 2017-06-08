package geneticalgorithm

import (
	"math/rand"
)

var (
	disc = NewStandardDiscretizer()
)

type Recombine func(parents <-chan DiscreteIndividuum, children chan<- DiscreteIndividuum)

func DummyRecombiner(parents <-chan DiscreteIndividuum, children chan<- DiscreteIndividuum) {
	for individuum := range parents {
		children <- individuum
	}
}

func OnePointCrossOver(parents <-chan DiscreteIndividuum, children chan<- DiscreteIndividuum) {

	last := <-parents
	children <- last
	for parent := range parents {
		parent1 := last
		parent2 := parent
		child := MakeIndividuum(parent1)
		for i := range child {
			mask := ((1<<disc.K) - 1) >> uint(rand.Intn(int(disc.K - 1)) + 1)
			child[i] = parent2[i] & mask | parent1[i] & (mask^0xFFFFFFFF)
		}
		children <- child
		last = parent
	}
}