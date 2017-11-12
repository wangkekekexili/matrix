package matrix

import (
	"fmt"
	"math"
)

type pivotPosition struct {
	row, col int
}

func (p pivotPosition) String() string {
	return fmt.Sprintf("(%v, %v)", p.row, p.col)
}

func RREF(m Matrix) {
	pivots := make([]*pivotPosition, 0)
	downward(m, 0, 0, &pivots)
	upward(m, pivots)
}

func downward(m Matrix, pRow, startCol int, pivots *[]*pivotPosition) {
	// Get next pivot position.
	pCol, ok := getNextPivotCol(m, pRow, startCol)
	if !ok {
		return
	}
	*pivots = append(*pivots, &pivotPosition{pRow, pCol})

	// Move row to get the maximum absolute pivot value.
	maxPivotValue(m, pRow, pCol)

	// Replace to make the elements below 0.
	clearDown(m, pRow, pCol)

	downward(m, pRow+1, pCol+1, pivots)
}

func getNextPivotCol(m Matrix, startRow, startCol int) (pivotCol int, ok bool) {
	numRow, numCol := m.Dims()
	if startRow >= numRow || startCol >= numCol {
		return 0, false
	}
	for c := startCol; c != numCol; c++ {
		for r := startRow; r != numRow; r++ {
			if m.At(r, c) != 0 {
				return c, true
			}
		}
	}
	return 0, false
}

func maxPivotValue(m Matrix, pRow, pCol int) {
	numRow, numCol := m.Dims()
	if pRow == numRow-1 {
		return
	}

	rowWithMax := pRow
	maxAbs := math.Abs(m.At(pRow, pCol))
	for r := pRow + 1; r != numRow; r++ {
		if math.Abs(m.At(r, pCol)) > maxAbs {
			maxAbs = math.Abs(m.At(r, pCol))
			rowWithMax = r
		}
	}
	if rowWithMax == pRow {
		return
	}

	// Interchange.
	for c := pCol; c != numCol; c++ {
		temp := m.At(pRow, c)
		m.Set(pRow, c, m.At(rowWithMax, c))
		m.Set(rowWithMax, c, temp)
	}
}

func upward(m Matrix, pivots []*pivotPosition) {
	_, numCol := m.Dims()
	for i := len(pivots) - 1; i >= 0; i-- {
		pRow, pCol := pivots[i].row, pivots[i].col

		// Multiple to get 1 as pivot element.
		v := m.At(pRow, pCol)
		m.Set(pRow, pCol, 1)
		if pCol != numCol-1 {
			for c := pCol + 1; c != numCol; c++ {
				m.Set(pRow, c, m.At(pRow, c)/v)
			}
		}

		if i == 0 {
			break
		}

		clearUp(m, pRow, pCol)
	}
}

func clearDown(m Matrix, pRow, pCol int) {
	numRow, _ := m.Dims()
	for r := pRow + 1; r != numRow; r++ {
		if m.At(r, pCol) == 0 {
			continue
		}
		replaceToClear(m, pRow, pCol, r)
	}
}

func clearUp(m Matrix, pRow, pCol int) {
	for r := pRow - 1; r >= 0; r-- {
		if m.At(r, pCol) == 0 {
			continue
		}
		replaceToClear(m, pRow, pCol, r)
	}
}

func replaceToClear(m Matrix, pRow, pCol, rowToClear int) {
	_, numCol := m.Dims()
	multiplicand := -m.At(rowToClear, pCol) / m.At(pRow, pCol)
	m.Set(rowToClear, pCol, 0)
	if pCol != numCol-1 {
		for c := pCol + 1; c != numCol; c++ {
			if m.At(pRow, c) == 0 {
				continue
			}
			toReplace := m.At(rowToClear, c) + m.At(pRow, c)*multiplicand
			if math.Abs(toReplace) < EPSILON {
				toReplace = 0
			}
			m.Set(rowToClear, c, toReplace)
		}
	}
}
