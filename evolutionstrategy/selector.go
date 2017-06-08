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

	keys := make([]float64, len(m))
	for key := range m {
		keys = append(keys, key)
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
