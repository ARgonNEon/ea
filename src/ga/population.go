package ga

import "fmt"
import "bytes"

type Population struct {
	individua []Individuum
	age       int
}

func GenerateStartPopulation(size int) Population {
	population := make([]Individuum, size)
	for i := range population {
		population[i] = makeRandomIndividuum()
	}
	return Population{population, 0}
}

func (pop Population) findBest() (best Individuum, index int) {
	minFitness := 1e9
	for i, individuum := range pop.individua {
		if fitness := individuum.getFitness(); fitness < minFitness {
			minFitness = fitness
			best = individuum
			index = i
		}
	}
	return
}

func (pop Population) getSize() int {
	return len(pop.individua)
}

func (pop *Population) exchange(index int, individuum Individuum) {
	pop.individua[index] = individuum
}

func (pop Population) String() string {
	var ss bytes.Buffer

	ss.WriteString("Population:\n")
	ss.WriteString(fmt.Sprintf("\t->Age: %d\n", pop.age))
	ss.WriteString(fmt.Sprintf("\t->Popsize: %d\n", len(pop.individua)))
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
