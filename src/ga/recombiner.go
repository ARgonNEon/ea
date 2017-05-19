package ga

type Recombine func(parents <-chan Individuum, children chan<- Individuum, popsize int)

func DummyRecombiner(parents <-chan Individuum, children chan<- Individuum, popsize int) {
	defer close(children)
	for i:=0; i<popsize; i++ {
		children <- <-parents
	}
}