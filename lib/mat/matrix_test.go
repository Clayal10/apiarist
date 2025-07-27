package mat_test

import (
	"testing"

	"github.com/Clayal10/mathGen/lib/mat"
)

func TestBasicMatrix(t *testing.T) {
	m1 := mat.NewMatrix([]float64{1.0, 2.0}, 1, 2)
	m2 := mat.NewMatrix([]float64{1.0, 2.0, 3.0, 4.0}, 2, 2)
	resultMatrix := mat.Mul(m1, m2)

	if resultMatrix.Values[0][0] != 7.0 ||
		resultMatrix.Values[0][1] != 10.0 {
		t.Fail()
	}
}
