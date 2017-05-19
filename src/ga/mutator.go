package ga

type Mutator interface {
	Mutate(<-chan Individuum, chan<- Individuum)
}
