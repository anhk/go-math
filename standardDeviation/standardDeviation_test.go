package standardDeviation

import (
	"fmt"
	"github.com/anhk/go-math/arithmeticMean"
	"testing"
)

var X = []float64{
	-0.60207, -0.85543, 0.4084, 0.60292,
	0.14554, -0.11812, -0.22425, -0.25985,
	-0.26579, 0.70213, 0.32766, 0.55666,
	0.50528, 0.62256, 0.21262, 0.36577,
	0.85273, 1.0159, 0.53494, 1.4023,
}

func TestStandardDeviation(t *testing.T) {
	m := arithmeticMean.NewArithmeticMean()
	m.Add(X...)
	mean := m.Result()

	fmt.Println(mean)

	v := NewStandardDeviation(mean)
	v.Add(X...)
	fmt.Println(v.Result())
}

func TestStandardDeviationResult(t *testing.T) {
	mean := arithmeticMean.ArithmeticMeanResult(X...)
	fmt.Println(StandardDeviationResult(mean, X...))
}
