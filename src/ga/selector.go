package ga

type OptimizeFunction func(Individuum) float64

type Selector interface {
	Select(<-chan Individuum, chan<- Individuum, OptimizeFunction) Population
}
