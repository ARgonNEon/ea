package ga

type Mutator interface {
	mutate(Individuum, int) Individuum
}
