package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	RowView            [][]int
	ColView            [][]int
	RowCount, ColCount int
}

func New(s string) (*Matrix, error) {
	errNonUniformMatrix := errors.New("non-uniform matrix")
	matrixLines := strings.Split(s, "\n")
	matrix := &Matrix{RowCount: len(matrixLines), ColCount: len(strings.Split(matrixLines[0], " "))}
	matrix.RowView = make([][]int, matrix.RowCount)
	matrix.ColView = make([][]int, matrix.ColCount)
	for rIdx, line := range matrixLines {
		nums := strings.Split(strings.TrimSpace(line), " ")
		if len(nums) != matrix.ColCount {
			return nil, errNonUniformMatrix
		}
		for cIdx := 0; cIdx < len(nums); cIdx++ {
			numVal, err := strconv.Atoi(nums[cIdx])
			if err != nil {
				return nil, err
			}
			matrix.RowView[rIdx] = append(matrix.RowView[rIdx], numVal)
			matrix.ColView[cIdx] = append(matrix.ColView[cIdx], numVal)
		}
	}
	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	columns := make([][]int, m.ColCount)
	for cIdx := 0; cIdx < m.ColCount; cIdx++ {
		columns[cIdx] = make([]int, m.RowCount)
		copy(columns[cIdx], m.ColView[cIdx])
	}
	return columns
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.RowCount)
	for rIdx := 0; rIdx < m.RowCount; rIdx++ {
		rows[rIdx] = make([]int, m.ColCount)
		copy(rows[rIdx], m.RowView[rIdx])
	}
	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= m.RowCount || col < 0 || col >= m.ColCount {
		return false
	}
	m.RowView[row][col] = val
	m.ColView[col][row] = val
	return true
}
