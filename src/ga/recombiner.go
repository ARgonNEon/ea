package ga

type Recombine func(parents <-chan Individuum, children chan<- Individuum, popsize int)

func DummyRecombiner(parents <-chan Individuum, children chan<- Individuum, popsize int) {

}