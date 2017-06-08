package evolutionstrategy

import (
	"fmt"

	"wesx.de/ArneS/ea/optimizer"
)

type Population struct {
	individuals []optimizer.Individuum
}

func MakePopulation(popsize int) Population {
	return Population{make([]optimizer.Individuum, popsize)}
}

func (p Population) streamIndividuals(out chan<- optimizer.Individuum) {
	defer close(out)
	for _, individuum := range p.individuals {
		out <- individuum
	}
}

func (p *Population) collectIndividuals(in <-chan optimizer.Individuum) {
	for i := 0; i < cap(p.individuals); i++ {
		p.individuals[i] = <-in
	}
}

func (p Population) findBest() optimizer.Individuum {
	bestValue := p.individuals[0].GetPhenotype()
	best := p.individuals[0]
	for _, individuum := range p.individuals {
		if v := individuum.GetPhenotype(); v < bestValue {
			best = individuum
		}
	}
	return best
}

func (p Population) String() string {
	return fmt.Sprintf("Size: %d, Best: %s", len(p.individuals), p.findBest().String())
}
