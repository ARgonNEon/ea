package evolutionstrategy

import "wesx.de/ArneS/ea/optimizer"

type Population struct {
	individuals []optimizer.Individuum
}

func MakePopulation(popsize int) Population {
	return Population{make([]optimizer.Individuum, popsize)}
}

func (p Population) streamIndividuals(out chan<- optimizer.Individuum) {
	for _, individuum := range p.individuals {
		out <- individuum
	}
}

func (p *Population) collectIndividuals(in <-chan optimizer.Individuum) {
	for i := 0; i < cap(p.individuals); i++ {
		p.individuals[i] = <-in
	}
}
