package optimizer

type IsOptimized func(individuum Fitnessable) bool

type Fitnessable interface {
	GetPhenotype() float64
}

type Optimizer interface {
	Optimize(optimized IsOptimized, verbose bool) Individuum
}
