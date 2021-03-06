package geneticalgorithm

import (
	"math/rand"
	"time"

	"github.com/aaarne/ea/ackley"
	"github.com/aaarne/ea/optimizer"
)

var discretize = NewStandardDiscretizer()

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DiscreteIndividuum []int

func MakeRandomIndividuum(nGenes int) DiscreteIndividuum {
	individuum := make([]int, nGenes)
	for i := range individuum {
		individuum[i] = rand.Intn(discretize.getMaxCode())
	}
	return individuum
}

func MakeIndividuum(template DiscreteIndividuum) DiscreteIndividuum {
	return make([]int, len(template))
}

func (individuum DiscreteIndividuum) GetPhenotype() float64 {
	var values []float64 = individuum.ToFloatIndividuum()
	return ackley.Ackley(values)
}

func (individuum DiscreteIndividuum) ToFloatIndividuum() optimizer.Individuum {
	values := make([]float64, 0, len(individuum)+1)
	for _, val := range individuum {
		values = append(values, discretize.Code2Value(val))
	}
	values = append(values, 0)
	return values
}

func (individuum DiscreteIndividuum) String() string {
	return individuum.ToFloatIndividuum().String()
}
