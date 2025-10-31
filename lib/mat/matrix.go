package mat

import (
	"math/rand/v2"
)

type Matrix struct {
	Width, Height int
	Values        [][]float64
}

func NewMatrix(input []float64, height, width int) *Matrix {
	m := newBlankMatrix(height, width)
	count := 0
	for i := range m.Values {
		for j := range m.Values[i] {
			m.Values[i][j] = input[count]
			count++
		}
	}
	return m
}
func Mul(one, two *Matrix, function func(float64) float64) *Matrix {
	if one.Width != two.Height {
		two = transpose(two)
	}

	m := newBlankMatrix(one.Height, two.Width)

	for i := range one.Height {
		for j := range two.Width {
			var sum float64 = 0
			for k := range one.Width {
				sum += one.Values[i][k] * two.Values[k][j]
			}
			if function != nil {
				sum = function(sum)
			}
			m.Values[i][j] = sum
			sum = 0
		}
	}
	return m
}

func NewRandomMatrixValues(height, width int) ([]float64, int, int) {
	list := make([]float64, height*width)
	for i := range list {
		list[i] = rand.Float64()
	}
	return list, height, width
}

// GetSize returns the number of values used in the matrix.
func (mat *Matrix) GetSize() int {
	return mat.Height * mat.Width
}

func (mat *Matrix) GetValueList() []float64 {
	list := make([]float64, mat.Height*mat.Width)
	index := 0
	for i := range mat.Height {
		for j := range mat.Width {
			list[index] = mat.Values[i][j]
			index++
		}
	}
	return list
}

func newBlankMatrix(height, width int) *Matrix {
	m := &Matrix{
		Width:  width,
		Height: height,
		Values: make([][]float64, height),
	}
	for i := range m.Values {
		m.Values[i] = make([]float64, width)
	}
	return m
}

func transpose(m *Matrix) *Matrix {
	v := make([]float64, m.Width*m.Height)
	count := 0
	for i := range m.Height {
		for j := range m.Width {
			v[count] = m.Values[i][j] // flips values
			count++
		}
	}
	return NewMatrix(v, m.Height, m.Width)
}
