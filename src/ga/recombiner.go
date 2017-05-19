package ga

type Recombiner interface {
	Recombine(Population) Population
}
