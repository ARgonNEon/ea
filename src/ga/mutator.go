package ga

type Mutate func(individuals <-chan Individuum, mutated chan<- Individuum)

func DummyMutator(individuals <-chan Individuum, mutated chan<- Individuum) {
	defer close(mutated)
	for individuum := range individuals {
		mutated <- individuum
	}
}