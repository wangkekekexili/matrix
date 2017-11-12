package matrix

import (
	"bytes"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		m   Matrix
		exp string
	}{
		{
			m:   mat.NewDense(1, 1, []float64{0}),
			exp: "0.0 \n",
		},
		{
			m:   mat.NewDense(2, 2, []float64{1.37, 2.2, 3.5, 4.2}),
			exp: "1.4 2.2 \n3.5 4.2 \n",
		},
	}

	for _, test := range tests {
		var b bytes.Buffer
		Print(&b, test.m, 1)
		if test.exp != b.String() {
			t.Fatalf("expected to get %v; got %v", test.exp, b.String())
		}
	}
}
