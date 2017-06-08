package evolutionstrategy

import (
	"math"
	"math/rand"

	"wesx.de/ArneS/ea/optimizer"
)

type Mutator struct {
	Popsize    int
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

func (m Mutator) Identity(in <-chan optimizer.Individuum, out chan<- optimizer.Individuum) {
	defer close(out)
	for individuum := range in {
		out <- individuum
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

func (m Mutator) adaptiveGaussianMutator(parent optimizer.Individuum) optimizer.Individuum {
	child := parent.CreateNew()
	sigma := parent[6] * math.Exp(rand.NormFloat64()/math.Sqrt(float64(m.Popsize)))
	for index, value := range parent[:6] {
		u := rand.NormFloat64() * parent[6]
		child[index] = value + u
		child[index] = math.Max(m.LowerBound, child[index])
		child[index] = math.Min(m.UpperBound, child[index])
	}
	child[6] = sigma
	return child
}
