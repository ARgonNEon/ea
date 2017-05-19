package ga

import "math/rand"

type Mutate func(individuals <-chan Individuum, mutated chan<- Individuum)

func DummyMutator(individuals <-chan Individuum, mutated chan<- Individuum) {
	defer close(mutated)
	for individuum := range individuals {
		mutated <- individuum
	}
}

func NonUniformMutator(individuals <-chan Individuum, mutated chan<- Individuum) {
	defer close(mutated)
	const sigma = 100
	for individuum := range individuals {
		for i := range individuum {
			if rand.Intn(2) == 0 {
				individuum[i] += int(rand.NormFloat64() * sigma)
			}
		}
		mutated <- individuum
	}
}