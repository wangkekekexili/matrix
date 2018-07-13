package matrix

import "fmt"

type Vector struct {
	values []float64
	err    error
}

func NewVector(values []float64) *Vector {
	v := &Vector{values: make([]float64, len(values))}
	copy(v.values, values)
	return v
}

func (v *Vector) At(i int) float64 {
	v.checkBound(i)
	return v.values[i]
}

func (v *Vector) checkBound(i int) {
	if i < 0 || i > v.Len() {
		v.err = fmt.Errorf("index %d out of bound", i)
	}
}

func (v *Vector) Err() error {
	return v.err
}

func (v *Vector) Len() int {
	return len(v.values)
}

func (v *Vector) Set(i int, f float64) {
	v.checkBound(i)
	v.values[i] = f
}
