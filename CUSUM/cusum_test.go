package CUSUM

// implement of https://en.wikipedia.org/wiki/CUSUM

import (
	"fmt"
	"testing"
)

var X = []float64{
	-0.60207, -0.85543, 0.4084, 0.60292,
	0.14554, -0.11812, -0.22425, -0.25985,
	-0.26579, 0.70213, 0.32766, 0.55666,
	0.50528, 0.62256, 0.21262, 0.36577,
	0.85273, 1.0159, 0.53494, 1.4023,
}

func TestCUSUM(t *testing.T) {
	mean := 0.0   // 算术平均数
	sigmaX := 0.5 // 标准差
	weight := 0.5 // 权重

	c := NewCUSUM(mean, sigmaX, weight)
	for _, v := range X {
		c.Add(v)
		r1, r2 := c.Result()
		fmt.Printf("%6f\t%6f\t%6f\n", v, r1, r2)
	}
}
