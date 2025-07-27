package mat

type Matrix struct {
	width, height int
	Values        [][]float64
}

func NewMatrix(input []float64, height, width int) Matrix {
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
func Mul(one, two Matrix) Matrix {
	if one.width != two.height {
		transpose(two)
	}

	m := newBlankMatrix(one.height, two.width)

	for i := range one.height {
		for j := range two.width {
			var sum float64 = 0
			for k := range one.width {
				sum += one.Values[i][k] * two.Values[k][j]
			}
			m.Values[i][j] = sum
			sum = 0
		}
	}
	return m
}
func newBlankMatrix(height, width int) Matrix {
	m := Matrix{
		width:  width,
		height: height,
		Values: make([][]float64, height),
	}
	for i := range m.Values {
		m.Values[i] = make([]float64, width)
	}
	return m
}

func transpose(m Matrix) Matrix {
	v := make([]float64, m.width*m.height)
	count := 0
	for i := range m.height {
		for j := range m.width {
			v[count] = m.Values[i][j] // flips values
			count++
		}
	}
	return NewMatrix(v, m.height, m.width)
}
