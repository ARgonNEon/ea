package ga

type OptimizeFunction func(Individuum) float64

type Select func(individuals <-chan Individuum, selected chan<- Individuum, optimizer OptimizeFunction)