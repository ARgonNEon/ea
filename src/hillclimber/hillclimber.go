package hillclimber

import "ackley"

type Hillclimber struct {
	nmax int
	delta float64
	dim int
}

func (hill Hillclimber) Climb(initialValue []float64) []float64 {
	bestpoint := make([]float64, hill.dim)
	copy(bestpoint, initialValue)
	bestval := ackley.Ackley(initialValue)
	for i:=0; i<hill.nmax; i++ {
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

func (hill Hillclimber) calculateNeighbors(point []float64, neighbors chan []float64) {
	for i:=0; i<hill.dim; i++ {
		for _, sign := range []float64{-1, 1} {
			neighbor := make([]float64, hill.dim)
			copy(neighbor, point)
			neighbor[i] += sign * hill.delta
			neighbors <- neighbor
		}
	}
	close(neighbors)
}