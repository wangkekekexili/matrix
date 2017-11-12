package matrix

import (
	"bytes"
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestRREF(t *testing.T) {
	tests := []struct {
		m, expRref Matrix
	}{
		{
			m:       mat.NewDense(1, 1, []float64{1}),
			expRref: mat.NewDense(1, 1, []float64{1}),
		},
		{
			m:       mat.NewDense(3, 6, []float64{0, 3, -6, 6, 4, -5, 3, -7, 8, -5, 8, 9, 3, -9, 12, -9, 6, 15}),
			expRref: mat.NewDense(3, 6, []float64{1, 0, -2, 3, 0, -24, 0, 1, -2, 2, 0, -7, 0, 0, 0, 0, 1, 4}),
		},
	}

	for _, test := range tests {
		RREF(test.m)
		if !SameValues(test.expRref, test.m) {
			var a, b bytes.Buffer
			Print(&a, test.m, 0)
			Print(&b, test.expRref, 0)
			t.Fatalf("expected to get\n%v;got\n%v", b.String(), a.String())
		}
	}
}

func TestDownward(t *testing.T) {
	tests := []struct {
		m, expRef Matrix
		expPivots []*pivotPosition
	}{
		{
			m:         mat.NewDense(1, 1, []float64{1}),
			expRef:    mat.NewDense(1, 1, []float64{1}),
			expPivots: []*pivotPosition{{0, 0}},
		},
		{
			m:         mat.NewDense(2, 2, []float64{1, 4, 2, 0}),
			expRef:    mat.NewDense(2, 2, []float64{2, 0, 0, 4}),
			expPivots: []*pivotPosition{{0, 0}, {1, 1}},
		},
	}

	for _, test := range tests {
		gotPivots := make([]*pivotPosition, 0)
		downward(test.m, 0, 0, &gotPivots)

		if !SameValues(test.expRef, test.m) {
			var a, b bytes.Buffer
			Print(&a, test.m, 0)
			Print(&b, test.expRef, 0)
			t.Fatalf("expected to get\n%v;got\n%v", b.String(), a.String())
		}

		if !reflect.DeepEqual(test.expPivots, gotPivots) {
			t.Fatalf("expected to get %v; got %v", test.expPivots, gotPivots)
		}
	}
}

func TestUpward(t *testing.T) {
	tests := []struct {
		ref, expRref   Matrix
		pivotPositions []*pivotPosition
	}{
		{
			ref:            mat.NewDense(1, 1, []float64{1}),
			pivotPositions: []*pivotPosition{{0, 0}},
			expRref:        mat.NewDense(1, 1, []float64{1}),
		},
		{
			ref:            mat.NewDense(3, 6, []float64{3, -9, 12, -9, 6, 15, 0, 2, -4, 4, 2, -6, 0, 0, 0, 0, 1, 4}),
			pivotPositions: []*pivotPosition{{0, 0}, {1, 1}, {2, 4}},
			expRref:        mat.NewDense(3, 6, []float64{1, 0, -2, 3, 0, -24, 0, 1, -2, 2, 0, -7, 0, 0, 0, 0, 1, 4}),
		},
	}

	for _, test := range tests {
		upward(test.ref, test.pivotPositions)
		if !SameValues(test.expRref, test.ref) {
			var a, b bytes.Buffer
			Print(&a, test.ref, 0)
			Print(&b, test.expRref, 0)
			t.Fatalf("expected to get\n%v;got\n%v", b.String(), a.String())
		}
	}
}
