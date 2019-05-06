package optimizer

import (
	"bytes"
	"fmt"

	"math/rand"

	"github.com/aaarne/ea/ackley"
)

type Individuum []float64

func (individuum Individuum) GetPhenotype() float64 {
	return ackley.Ackley([]float64(individuum[:6]))
}

func (individuum Individuum) createFloatString() string {
	var ss bytes.Buffer

	ss.WriteString("[")
	for _, val := range individuum[:6] {
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

func (individuum Individuum) CreateNew() Individuum {
	return MakeFromTemplate(individuum)
}

func (individuum Individuum) String() string {
	return fmt.Sprintf("Individuum: [Phenotype: %.6f, Genotype %v, Sigma: %.6f]", individuum.GetPhenotype(), individuum.createFloatString(), individuum[6])
}
