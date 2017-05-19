package ga

type Recombiner interface {
	Recombine(parents <-chan Individuum, children chan<- Individuum) Population
}
