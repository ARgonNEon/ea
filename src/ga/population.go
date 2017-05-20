package ga

import "fmt"
import "bytes"

type Population struct {
	individuals []Individuum
	age         int
}

func GenerateStartPopulation(size int) Population {
	population := GenerateEmtpyPopulation(size)
	ch := make(chan Individuum, 10)
	go func() {
		for i:=0; i<size; i++ {
			ch <- MakeRandomIndividuum(NGenes)
		}
	}()
	population.collectIndividuals(ch)
	return population
}

func GenerateEmtpyPopulation(size int) Population {
	population := make([]Individuum, size)
	return Population{population, 0}
}

func (pop *Population) collectIndividuals(chIndivuduals <-chan Individuum) {
	for i:=0; i<cap(pop.individuals); i++ {
		pop.individuals[i] = <-chIndivuduals
	}
	pop.age++
}

func (pop Population) streamIndividuals(chIndividuals chan<-Individuum) {
	for _, individuum := range pop.individuals {
		chIndividuals <- individuum
	}

}

func (pop Population) findBest() (best Individuum, index int) {
	minFitness := 1e9
	for i, individuum := range pop.individuals {
		if fitness := individuum.getFitness(); fitness < minFitness {
			minFitness = fitness
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
	var ss bytes.Buffer

	ss.WriteString("Population:\n")
	ss.WriteString(fmt.Sprintf("\t->Age: %d\n", pop.age))
	ss.WriteString(fmt.Sprintf("\t->Popsize: %d\n", len(pop.individuals)))
	ss.WriteString("\t->Best Individuum:\n")
	best, index := pop.findBest()
	ss.WriteString(fmt.Sprintf("\t\t->Index: %d\n", index))
	ss.WriteString(fmt.Sprintf("\t\t->Phenotype: %.3f\n", best.getFitness()))
	ss.WriteString(fmt.Sprintf("\t\t->Genotype: "))
	ss.WriteString(fmt.Sprintln(best))
	return ss.String()
}

func (pop Population) Analyze() {
	fmt.Print(pop)
}
