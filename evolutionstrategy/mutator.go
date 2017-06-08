package evolutionstrategy

import (
	"math"
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
		for i := 0; i < m.Lambda; i++ {
			out <- m.gaussianMutator(individuum)
		}
	}
}

func (m Mutator) gaussianMutator(parent optimizer.Individuum) optimizer.Individuum {
	child := parent.CreateNew()
	for index, value := range parent {
		u := rand.NormFloat64() * m.Sigma
		child[index] = value + u
		child[index] = math.Max(m.LowerBound, child[index])
		child[index] = math.Min(m.UpperBound, child[index])
	}
	return child
}
