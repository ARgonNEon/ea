package ga

type Recombine func(parents <-chan Individuum, children chan<- Individuum)