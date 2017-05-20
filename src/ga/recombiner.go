package ga

import (
	"math/rand"
)

type Recombine func(parents <-chan Individuum, children chan<- Individuum)

func DummyRecombiner(parents <-chan Individuum, children chan<- Individuum) {
	for individuum := range parents {
		children <- individuum
	}
}

func OnePointCrossOver(parents <-chan Individuum, children chan<- Individuum) {

	last := <-parents
	children <- last
	for parent := range parents {
		parent1 := last
		parent2 := parent
		child := MakeIndividuum(parent1)
		for i := range child {
			mask := 0x1FFFF >> uint(rand.Intn(16) + 1)
			child[i] = parent2[i] & mask | parent1[i] & (mask^0xFFFFFFFF)
		}
		children <- child
		last = parent
	}
}