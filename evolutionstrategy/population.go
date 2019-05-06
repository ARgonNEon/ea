package evolutionstrategy

import (
	"fmt"

	"github.com/aaarne/ea/optimizer"
)

type Population struct {
	individuals []optimizer.Individuum
	age         int
}

func MakePopulation(popsize int) Population {
	return Population{
		individuals: make([]optimizer.Individuum, popsize),
		age:         0,
	}
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
	p.age++
}

func (p Population) matchAny(optimized optimizer.IsOptimized) bool {
	for _, individuum := range p.individuals {
		if optimized(individuum) {
			return true
		}
	}
	return false
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
	return fmt.Sprintf("Size: %d, Best: %s, Age: %d", len(p.individuals), p.findBest().String(), p.age)
}
