package matrix

import (
	"fmt"
	"io"
	"strconv"
)

func Print(w io.Writer, m Matrix, p int) {
	rows, cols := m.Dims()
	colMaxLength := make([]int, cols)
	content := make([][]string, rows)

	for row := 0; row != rows; row++ {
		for col := 0; col != cols; col++ {
			str := strconv.FormatFloat(m.At(row, col), 'f', p, 64)

			if content[row] == nil {
				content[row] = make([]string, cols)
			}
			content[row][col] = str

			if colMaxLength[col] < len(str) {
				colMaxLength[col] = len(str)
			}
		}
	}

	colToFormat := make(map[int]string, cols)
	for col, maxLength := range colMaxLength {
		colToFormat[col] = fmt.Sprintf("%%%v.%df", maxLength+1, p)
	}

	for row := 0; row != rows; row++ {
		for col := 0; col != cols; col++ {
			fmt.Fprintf(w, colToFormat[col], m.At(row, col))
		}
		fmt.Fprintln(w)
	}
}
