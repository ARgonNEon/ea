package evolutionstrategy

import (
	"math/rand"

	"wesx.de/ArneS/ea/optimizer"
)

type Mutator struct {
	Lambda     int
	Sigma      float64
	UpperBound float64
	LowerBound float64
}

func (m Mutator) Mutate(in <-chan optimizer.Individuum, out chan<- optimizer.Individuum) {
	defer close(out)
	for individuum := range in {

	}
}

func (m Mutator) gaussianMutator(parent optimizer.Individuum) optimizer.Individuum {
	for index, value := range parent {
		u := rand.NormFloat64() * m.Sigma

	}
}
