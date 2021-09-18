package variance

import "math"

/**
 * 方差
 **/

type Variance struct {
	mean   float64
	result float64
	count  int64
}

// NewVariance
// - mean: 算数平均数
func NewVariance(mean float64) *Variance {
	return &Variance{mean: mean}
}

func (v *Variance) Add(x ...float64) {
	for _, d := range x {
		v.result += math.Pow(d-v.mean, 2)
		v.count++
	}
}

func (v *Variance) Result() float64 {
	if v.count == 0 {
		return 0
	}
	return v.result / float64(v.count)
}

// VarianceResult 直接计算方差
// - mean: 算术平均数
func VarianceResult(mean float64, x ...float64) float64 {
	s := NewVariance(mean)
	s.Add(x...)
	return s.Result()
}
