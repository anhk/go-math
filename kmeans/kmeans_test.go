package kmeans

import (
	"fmt"
	"testing"
)

var (
	P = []Point{{1.5, 2}, {1, 3}, {2, 4}, {9, 19}}
	Q = []Point{{1, 2}, {1, 3}, {1, 4}, {1, 19}}
)

func TestKMeans(t *testing.T) {
	k := NewKMeans(4)
	c := k.Run(P)
	for _, cc := range c {
		fmt.Println(cc.Centroide, cc.Points)
	}

	fmt.Println("++++")
	c2 := k.Run(Q)
	for _, cc := range c2 {
		fmt.Println(cc.Centroide, cc.Points)
	}
}
