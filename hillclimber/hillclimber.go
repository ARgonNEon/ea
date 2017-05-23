package hillclimber

import "wesx.de/ArneS/ea/ackley"

type Hillclimber struct {
	Nmax  int
	Delta float64
	Dim   int
}

func (hill Hillclimber) Climb(initialValue []float64) []float64 {
	bestpoint := make([]float64, hill.Dim)
	copy(bestpoint, initialValue)
	bestval := ackley.Ackley(initialValue)
	for i:=0; i<hill.Nmax; i++ {
		c := make(chan []float64, 5)
		go hill.calculateNeighbors(bestpoint, c)
		for neighbor := range c {
			if ackley.Ackley(neighbor) < bestval {
				copy(bestpoint, neighbor)
				bestval = ackley.Ackley(bestpoint)
			}
		}
	}
	return bestpoint
}

func (hill Hillclimber) calculateNeighbors(point []float64, neighbors chan<- []float64) {
	defer close(neighbors)
	for i:=0; i<hill.Dim; i++ {
		for _, sign := range []float64{-1, 1} {
			neighbor := make([]float64, hill.Dim)
			copy(neighbor, point)
			neighbor[i] += sign * hill.Delta
			neighbors <- neighbor
		}
	}
}
