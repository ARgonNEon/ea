package ga

type OptimizeFunction func(Individuum) float64

type Selector interface {
	Select(individuals <-chan Individuum, selected chan<- Individuum, optimizer OptimizeFunction) Population
}
