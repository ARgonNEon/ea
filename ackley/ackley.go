package ackley

import "math"

type inputSlice []float64

const a = 20
const b = 0.2
const c = 2 * math.Pi

func Ackley(input inputSlice) float64 {
	dim := len(input)
	var accumulator float64
	accumulator += -a * math.Exp(-b*math.Sqrt((1.0/float64(dim))*input.squaredSum()))
	cosines := make(inputSlice, dim)
	for i, val := range input {
		cosines[i] = math.Cos(c * val)
	}
	accumulator += -math.Exp((1 / float64(dim)) * cosines.sum())
	accumulator += a
	accumulator += math.E
	return accumulator
}

func (input inputSlice) sum() float64 {
	var sum float64
	for _, val := range input {
		sum += val
	}
	return sum
}

func (input inputSlice) squaredSum() float64 {
	var sum float64
	for _, val := range input {
		sum += math.Pow(val, 2)
	}
	return sum
}
