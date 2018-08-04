package laff

import "math"

type Vector struct {
	values []float64
	err    error
}

func NewVector(values []float64) *Vector {
	v := &Vector{values: make([]float64, len(values))}
	copy(v.values, values)
	return v
}

func NewZeroVector(length int) *Vector {
	return &Vector{values: make([]float64, length)}
}

func (v *Vector) Size() int {
	return len(v.values)
}

func (v *Vector) At(i int) float64 {
	return v.values[i]
}

func (v *Vector) Set(i int, f float64) *Vector {
	v.values[i] = f
	return v
}

func (v *Vector) Clone() *Vector {
	w := NewVector(v.values)
	w.err = v.err
	return w
}

func (v *Vector) ForEach(fn func(int, float64) float64) *Vector {
	for i, e := range v.values {
		v.values[i] = fn(i, e)
	}
	return v
}

// Scale scales the vector in place.
func (v *Vector) Scale(f float64) *Vector {
	return v.ForEach(func(_ int, v float64) float64 { return f * v })
}

// AXPY operation. v is y.
func (v *Vector) AXPY(alpha float64, x *Vector) *Vector {
	if v.Size() != x.Size() {
		v.err = ErrVectorSizeMismatch{
			left:  v.Size(),
			right: x.Size(),
		}
		return v
	}

	v.ForEach(func(i int, f float64) float64 {
		return alpha*x.values[i] + f
	})

	return v
}

func (v *Vector) Dot(right *Vector) float64 {
	var result float64
	for i := range v.values {
		result += v.values[i] * right.values[i]
	}
	return result
}

func (v *Vector) DotSafe(right *Vector) (float64, error) {
	if v.Size() != right.Size() {
		return 0, ErrVectorSizeMismatch{
			left:  v.Size(),
			right: right.Size(),
		}
	}
	return v.Dot(right), nil
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}
