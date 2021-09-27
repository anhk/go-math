package kmeans

import (
	"fmt"
	"testing"
)

var (
	P = []Point{
		{[]float64{1.5, 2, 7}},
		{[]float64{1, 3, 0}},
		{[]float64{2, 4, -1}},
		{[]float64{9, 19, 17}},
	}
)

func TestKMeans(t *testing.T) {
	k := NewKMeans(4)
	c := k.Run(P)
	for _, cc := range c {
		fmt.Println(cc.Centroide, cc.Points)
	}
}
