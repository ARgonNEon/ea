package optimizer

type IsOptimized func(individuum Fitnessable) bool

type Fitnessable interface {
	GetPhenotype() float64
}
