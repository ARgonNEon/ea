package geneticalgorithm

import (
	"fmt"
	"bytes"
	"math/rand"
	"time"
	"ackley"
)

var discretize = NewStandardDiscretizer()

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Fitness interface {
	getFitness() float64
}

type Individuum []int
type floatIndividuum []float64

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

func (individuum Individuum) GetFitness() float64 {
	var values []float64 = individuum.ToFloatIndividuum()
	return ackley.Ackley(values)
}

func (individuum Individuum) ToFloatIndividuum() floatIndividuum {
	values := make([]float64, 0, len(individuum))
	for _, val := range individuum {
		values = append(values, discretize.Code2Value(val))
	}
	return values
}

func (fIndividuum floatIndividuum) String() string {
	var ss bytes.Buffer

	ss.WriteString("[")
	for _, val := range fIndividuum {
		ss.WriteString(fmt.Sprintf("%5f ", val))
	}
	ss.WriteString("\b]")
	return ss.String()
}

func (individuum Individuum) String() string {
	return fmt.Sprintf("Individuum: [Phenotype: %.6f, Genotype %v]", individuum.GetFitness(), individuum.ToFloatIndividuum())
}
