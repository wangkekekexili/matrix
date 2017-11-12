package matrix

import (
	"math"
)

var EPSILON = 0.00000001

type Matrix interface {
	Dims() (r, c int)
	At(i, j int) float64
	Set(i, j int, v float64)
}

// SameValues returns true if a and b are of the same size and contain same value at the same position.
func SameValues(a, b Matrix) bool {
	aRow, aCol := a.Dims()
	bRow, bCol := b.Dims()
	if aRow != bRow || aCol != bCol {
		return false
	}

	for r := 0; r != aRow; r++ {
		for c := 0; c != aCol; c++ {
			if math.Abs(a.At(r, c)-b.At(r, c)) > EPSILON {
				return false
			}
		}
	}

	return true
}
