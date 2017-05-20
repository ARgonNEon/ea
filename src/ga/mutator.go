package ga

import "math/rand"

type MutateContext struct {
	Age int
	MaxAge int
}

type Mutate func(individuals <-chan Individuum, mutated chan<- Individuum, context MutateContext)

func DummyMutator(individuals <-chan Individuum, mutated chan<- Individuum, context MutateContext) {
	for individuum := range individuals {
		mutated <- individuum
	}
}

func nonUniformMutator(individuals <-chan Individuum, mutated chan<- Individuum, sigma float64) {
	for individuum := range individuals {
		for i := range individuum {
			if rand.Intn(2) == 0 {
				individuum[i] += int(rand.NormFloat64() * sigma)
			}
		}
		mutated <- individuum
	}
}

func NonUniformMutator(individuals <-chan Individuum, mutated chan<- Individuum, context MutateContext) {
	nonUniformMutator(individuals, mutated, 100)
}

func AdaptiveGaussianMutator(individuals <-chan Individuum, mutated chan<- Individuum, context MutateContext) {
	nonUniformMutator(individuals, mutated, 10 + 1500 * (1 - float64(context.Age)/float64(context.MaxAge)))
}