package ga

type OptimizeFunction func(Individuum) float64

type Selector interface {
	Select(Population, OptimizeFunction) Population
}
