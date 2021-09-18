package standardDeviation

import (
	"github.com/anhk/go-math/variance"
	"math"
)

/**
 * 标准差
 **/

type StandardDeviation struct {
	v *variance.Variance
}

// NewStandardDeviation
// - mean: 算数平均数
func NewStandardDeviation(mean float64) *StandardDeviation {
	return &StandardDeviation{v: variance.NewVariance(mean)}
}

func (s *StandardDeviation) Add(x ...float64) {
	s.v.Add(x...)
}

func (s *StandardDeviation) Result() float64 {
	return math.Sqrt(s.v.Result())
}

// StandardDeviationResult 直接计算标准差
// - mean: 算术平均数
func StandardDeviationResult(mean float64, x ...float64) float64 {
	s := NewStandardDeviation(mean)
	s.Add(x...)
	return s.Result()
}
