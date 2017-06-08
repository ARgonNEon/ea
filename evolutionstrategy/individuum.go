package evolutionstrategy

import (
	"bytes"
	"fmt"

	"wesx.de/ArneS/ea/ackley"
)

type Fitness interface {
	GetPhenotype() float64
}

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

func (individuum Individuum) String() string {
	return fmt.Sprintf("Individuum: [Phenotype: %.6f, Genotype %v]", individuum.GetPhenotype(), individuum.createFloatString())
}
