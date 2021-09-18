package CUSUM

/**
 * 累积和
 */

type CUSUM struct {
	mean, sigma, weight float64
	resultH, resultL    float64
}

// NewCUSUM
// - mean: 算术平均数
// - sigma: 标准差
// - weight: 权重，误差
func NewCUSUM(mean, sigma, weight float64) *CUSUM {
	if sigma == 0 {
		return nil
	}
	return &CUSUM{mean: mean, sigma: sigma, weight: weight}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (c *CUSUM) Add(x ...float64) {
	for _, d := range x {
		z := (d - c.mean) / c.sigma
		c.resultH = max(0, c.resultH+z-c.weight)
		c.resultL = max(0, c.resultL-z-c.weight)
	}
}

func (c *CUSUM) Result() (float64, float64) {
	return c.resultH, c.resultL
}
