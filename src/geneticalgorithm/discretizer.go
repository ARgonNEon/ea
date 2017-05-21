package geneticalgorithm

import (
	"math"
)

type Discretize interface {
	Code2Value(code int) float64
	Value2Code(value float64) int
}

type LinearDiscretizer struct {
	XMax    float64
	XMin    float64
	Epsilon float64
	K       int
}

func NewDiscretizer(xmin, xmax, epsilon float64) *LinearDiscretizer {
	k := int(math.Ceil(math.Log2((xmax - xmin) / epsilon)))
	return &LinearDiscretizer{xmax, xmin, epsilon, k}
}

func NewStandardDiscretizer() *LinearDiscretizer {
	return NewDiscretizer(XMin, XMax, Epsilon)
}

func (disc LinearDiscretizer) Code2Value(code int) float64 {
	fCode := float64(code)
	return fCode/(math.Pow(2, float64(disc.K))-1)*(disc.XMax-disc.XMin) + disc.XMin
}

func (disc LinearDiscretizer) Value2Code(value float64) int {
	return int(math.Floor((value-disc.XMin)/(disc.XMax-disc.XMin)*math.Pow(2, float64(disc.K)) - 1))
}

func (disc LinearDiscretizer) Value2RelativeCode(value float64) int {
	return int(math.Floor(value/(disc.XMax-disc.XMin)*math.Pow(2, float64(disc.K)) - 1))
}

func (disc LinearDiscretizer) getMaxCode() int {
	return disc.Value2Code(XMax)
}
