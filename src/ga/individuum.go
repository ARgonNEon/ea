package ga

import "math/rand"
import "time"
import "ackley"

var discretize = NewStandardDiscretizer()

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Fitness interface {
	getFitness() float64
}

type Individuum []int

func MakeRandomIndividuum(nGenes int) Individuum {
	individuum := make([]int, nGenes)
	for i := range individuum {
		individuum[i] = rand.Intn(discretize.getMaxCode())
	}
	return individuum
}

func MakeIndividuum(template Individuum) Individuum {
	return make([]int, len(template))
}

func (individuum Individuum) getFitness() float64 {
	values := make([]float64, 0, len(individuum))
	for _, val := range individuum {
		values = append(values, discretize.Code2Value(val))
	}
	return ackley.Ackley(values)
}
