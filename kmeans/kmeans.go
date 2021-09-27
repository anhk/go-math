package kmeans

import (
	"math"
	"math/rand"
)

/**
 * 聚类算法-kmeans with Distance(Euclidean)
 */

/**
 * 点
 */
type Point struct {
	M []float64
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func (p Point) DistanceTo(t Point) float64 {
	if len(t.M) != len(p.M) {
		return 0
	}
	var sum float64
	for i := range t.M {
		sum += math.Pow(t.M[i]-p.M[i], 2)
	}
	return math.Sqrt(sum)
}

func (p Point) Equals(point Point) bool {
	if len(point.M) != len(p.M) {
		return false
	}
	for i := range p.M {
		if !almostEqual(p.M[i], point.M[i]) {
			return false
		}
	}
	return true
}

// Add
func (p Point) Add(v Point) {
	for i := range v.M {
		p.M[i] += v.M[i]
	}
}

func (p Point) Div(d int) {
	for i := range p.M {
		p.M[i] /= float64(d)
	}
}

func (p Point) calcMax(v Point) {
	for i := range p.M {
		p.M[i] = math.Max(p.M[i], v.M[i])
	}
}

func (p Point) calcMin(v Point) {
	for i := range p.M {
		p.M[i] = math.Min(p.M[i], v.M[i])
	}
}

type KMeans struct {
	k        int
	clusters []Cluster
}

// 聚类结果
type Cluster struct {
	Centroide Point   // 中心点
	Points    []Point // 聚类内的点

	isConvergence bool // 是否完成聚类收敛
}

func (c *Cluster) addPoint(v Point) {
	c.Points = append(c.Points, v)
}

// calcNewCentroide: 求均值，作为新的质心点
func (c *Cluster) calcNewCentroide() Point {
	newCentroide := newZeroPoint(len(c.Centroide.M))
	for _, v := range c.Points {
		newCentroide.Add(v)
	}
	newCentroide.Div(len(c.Points))
	return newCentroide
}

func NewKMeans(k int) *KMeans {
	return &KMeans{k: k}
}

// initCentroides: 初始化k个质心点
func (k *KMeans) initCentroides(points []Point) {
	k.clusters = make([]Cluster, k.k)
	dimension := len(points[0].M)

	max := newMinPoint(dimension)
	min := newMaxPoint(dimension)

	for _, v := range points {
		max.calcMax(v)
		min.calcMin(v)
	}
	for i := 0; i < k.k; i++ {
		k.clusters[i].Centroide = randomBetween(min, max)
	}
}

func randomBetween(min Point, max Point) Point {
	p := Point{M: make([]float64, len(min.M))}
	for i := range min.M {
		p.M[i] = rand.Float64()*(max.M[i]-min.M[i]) + min.M[i]
	}
	return p
}

func newZeroPoint(dimension int) Point {
	return Point{M: make([]float64, dimension)}
}

func newMaxPoint(dimension int) Point {
	p := Point{M: make([]float64, dimension)}
	for i := range p.M {
		p.M[i] = math.MaxFloat64
	}
	return p
}

func newMinPoint(dimension int) Point {
	p := Point{M: make([]float64, dimension)}
	for i := range p.M {
		p.M[i] = math.SmallestNonzeroFloat64
	}
	return p
}

// isConvergence: 是否收敛
func (k *KMeans) isConvergence() bool {
	for _, v := range k.clusters {
		if !v.isConvergence {
			return false
		}
	}
	return true
}

func (k *KMeans) Run(points []Point) []Cluster {
	if len(points) == 0 {
		return nil
	}
	// 检查所有点维度是否相同
	dimension := len(points[0].M)
	for _, v := range points {
		if dimension != len(v.M) {
			return nil
		}
	}

	k.initCentroides(points)
	for !k.isConvergence() {
		k.initCluster()          // 重置分类中的点
		k.classifyPoint(points)  // 计算距离进行分类
		k.recalcularCentroides() // 重新计算质心点
	}

	return k.clusters
}

// 重置分类中的点
func (k *KMeans) initCluster() {
	for i := 0; i < k.k; i++ {
		k.clusters[i].Points = nil
	}
}

// classifyPoint: 计算距离进行分类
func (k *KMeans) classifyPoint(points []Point) {
	for _, v := range points {
		masCercano := 0
		minDistancia := math.MaxFloat64
		for i := 0; i < k.k; i++ {
			distancia := v.DistanceTo(k.clusters[i].Centroide)
			if minDistancia > distancia {
				minDistancia = distancia
				masCercano = i
			}
		}
		k.clusters[masCercano].addPoint(v)
	}
}

// recalcularCentroides: 重新计算质心点
func (k *KMeans) recalcularCentroides() {
	for i := 0; i < k.k; i++ {
		if len(k.clusters[i].Points) == 0 {
			k.clusters[i].isConvergence = true
			continue
		}

		// 求均值，作为新的质心点
		newCentroide := k.clusters[i].calcNewCentroide()
		if k.clusters[i].Centroide.Equals(newCentroide) {
			k.clusters[i].isConvergence = true
		} else {
			k.clusters[i].Centroide = newCentroide
		}
	}
}
