package evolutionstrategy

import (
	"math"
	"math/rand"
	"sort"

	"wesx.de/ArneS/ea/optimizer"
)

type Selector struct {
	Popsize int
}

func (s Selector) Select(in <-chan optimizer.Individuum, out chan<- optimizer.Individuum) {
	m := make(map[float64]optimizer.Individuum)

	for individuum := range in {
		m[math.Exp(-individuum.GetPhenotype())] = individuum
	}

	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	for i := 0; i < s.Popsize; i++ {
		for index := 0; index < len(m); index++ {
			if r := rand.Float64(); r < (1 / math.Pow(2, float64(index))) {
				out <- m[keys[index]]
			}
		}
	}
}

func (s Selector) SelectSorted(in <-chan optimizer.Individuum, out chan<- optimizer.Individuum) {
	m := make(map[float64]optimizer.Individuum)

	for individuum := range in {
		m[math.Exp(-individuum.GetPhenotype())] = individuum
	}

	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	for i := 1; i <= s.Popsize; i++ {
		out <- m[keys[len(keys)-i]]
	}
}

func (s Selector) SelectDynamically(in <-chan optimizer.Individuum, out chan<- optimizer.Individuum) {
	var input []optimizer.Individuum
	var totalAmount int
	for i := range in {
		input = append(input, i)
	}
	var sum float64
	distribution := make([]float64, len(input))
	for index, individuum := range input {
		value := -math.Exp(-individuum.GetPhenotype())
		sum += value
		distribution[index] = value
	}
	for index := range distribution {
		distribution[index] *= 1 / sum
	}

	for i, individuum := range input {
		amount := int(math.Floor(distribution[i] * float64(len(input))))
		for j := 0; j < amount; j++ {
			out <- individuum
			totalAmount++
		}
	}

	for i := 0; i < (len(input) - totalAmount); i++ {
		out <- input[rand.Intn(len(input))]
	}

}

func (s Selector) Identity(in <-chan optimizer.Individuum, out chan<- optimizer.Individuum) {
	for individuum := range in {
		out <- individuum
	}
}
