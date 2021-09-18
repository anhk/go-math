package arithmeticMean

/**
 * 算术平均数
 **/

type ArithmeticMean struct {
	summary float64
	count   int64
}

// NewArithmeticMean
func NewArithmeticMean() *ArithmeticMean {
	return &ArithmeticMean{}
}

func (m *ArithmeticMean) Add(x ...float64) {
	for _, d := range x {
		m.summary += d
		m.count++
	}
}

func (m *ArithmeticMean) Result() float64 {
	if m.count == 0 {
		return 0
	}
	return m.summary / float64(m.count)
}

// ArithmeticMeanResult 直接计算-算术平均数
func ArithmeticMeanResult(x ...float64) float64 {
	m := NewArithmeticMean()
	m.Add(x...)
	return m.Result()
}
