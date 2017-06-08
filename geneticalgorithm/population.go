package geneticalgorithm

import (
	"fmt"
)

type Population struct {
	individuals []DiscreteIndividuum
	age         uint64
}

func GenerateStartPopulation(size int) Population {
	population := GenerateEmtpyPopulation(size)
	ch := make(chan DiscreteIndividuum, 10)
	go func() {
		for i := 0; i < size; i++ {
			ch <- MakeRandomIndividuum(NGenes)
		}
	}()
	population.collectIndividuals(ch)
	return population
}

func GenerateEmtpyPopulation(size int) Population {
	population := make([]DiscreteIndividuum, size)
	return Population{population, 0}
}

func (pop *Population) collectIndividuals(chIndivuduals <-chan DiscreteIndividuum) {
	for i := 0; i < cap(pop.individuals); i++ {
		pop.individuals[i] = <-chIndivuduals
	}
	pop.age++
}

func (pop Population) streamIndividuals(chIndividuals chan<- DiscreteIndividuum) {
	for _, individuum := range pop.individuals {
		chIndividuals <- individuum
	}

}

func (pop Population) findBest() (best DiscreteIndividuum, index int) {
	minPhenotype := 1e9
	for i, individuum := range pop.individuals {
		if phenotype := individuum.GetPhenotype(); phenotype < minPhenotype {
			minPhenotype = phenotype
			best = individuum
			index = i
		}
	}
	return
}

func (pop Population) GetSize() int {
	return len(pop.individuals)
}

func (pop *Population) increaseAge() {
	pop.age++
}

func (pop Population) String() string {
	best, index := pop.findBest()

	return fmt.Sprintf("Population: [Age: %d, Popsize: %d, Best Individuum(index=%d): %v]",
		pop.age, len(pop.individuals), index, best)
}
