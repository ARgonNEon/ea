package optimizer

import (
	"bytes"
	"fmt"

	"math/rand"

	"wesx.de/ArneS/ea/ackley"
)

type Individuum []float64

func (individuum Individuum) GetPhenotype() float64 {
	return ackley.Ackley([]float64(individuum))
}

func (individuum Individuum) createFloatString() string {
	var ss bytes.Buffer

	ss.WriteString("[")
	for _, val := range individuum {
		ss.WriteString(fmt.Sprintf("%5f ", val))
	}
	ss.WriteString("\b]")
	return ss.String()
}

func MakeRandomIndividuum(nGenes int, max, min float64) Individuum {
	individuum := make(Individuum, nGenes)
	for index := range individuum {
		r := rand.Float64()*(max-min) + min
		individuum[index] = r
	}
	return individuum
}

func MakeFromTemplate(template Individuum) Individuum {
	individuum := make(Individuum, cap(template))
	return individuum
}

func (individuum Individuum) String() string {
	return fmt.Sprintf("Individuum: [Phenotype: %.6f, Genotype %v]", individuum.GetPhenotype(), individuum.createFloatString())
}
