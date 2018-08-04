package laff

import (
	"reflect"
	"testing"
)

func TestVector_Clone(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	w := v.Clone()
	if len(w.values) != 3 {
		t.Fatalf("w.values=%v;want %v", len(w.values), 3)
	}

	// Changing the values in the old vector won't change the cloned vector;
	v.values[0] = 10
	if w.values[0] != 1 {
		t.Fatalf("w.values[0]=%v;want %v", w.values[0], 1)
	}
}

func TestVector_Scale(t *testing.T) {
	v := NewVector([]float64{11, 2, 7})
	tests := []struct {
		factor float64
		exp    []float64
	}{
		{0, []float64{0, 0, 0}},
		{1, []float64{11, 2, 7}},
		{1.5, []float64{16.5, 3, 10.5}},
	}
	for _, test := range tests {
		got := v.Clone().Scale(test.factor)
		if !reflect.DeepEqual(got.values, test.exp) {
			t.Fatalf("got %v; want %v", got.values, test.exp)
		}
	}
}

func TestVector_AXPY(t *testing.T) {
	alpha := 3.0
	x := NewVector([]float64{1, 2, 3})
	y := NewVector([]float64{7, 8, 9})
	y.AXPY(alpha, x)

	if y.err != nil {
		t.Fatal(y.err)
	}
	exp := []float64{10, 14, 18}
	if !reflect.DeepEqual(y.values, exp) {
		t.Fatalf("got %v; want %v", y.values, exp)
	}
}
