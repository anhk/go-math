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
	X, Y float64
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func (p Point) DistanceTo(t Point) float64 {
	return math.Sqrt(math.Pow(t.X-p.X, 2) + math.Pow(t.Y-p.Y, 2))
}

func (p Point) Equals(point Point) bool {
	return almostEqual(p.X, point.X) && almostEqual(p.Y, point.Y)
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
	var x, y, sumX, sumY float64
	for _, v := range c.Points {
		sumX += v.X
		sumY += v.Y
	}
	x = sumX / float64(len(c.Points))
	y = sumY / float64(len(c.Points))
	return Point{X: x, Y: y}
}

func NewKMeans(k int) *KMeans {
	return &KMeans{k: k}
}

// initCentroides: 初始化k个质心点
func (k *KMeans) initCentroides(points []Point) {
	k.clusters = make([]Cluster, k.k)

	// 找到数据集的范围
	maxX := math.SmallestNonzeroFloat64
	maxY := math.SmallestNonzeroFloat64
	minX := math.MaxFloat64
	minY := math.MaxFloat64

	for _, v := range points {
		maxX = math.Max(maxX, v.X)
		maxY = math.Max(maxY, v.Y)
		minX = math.Min(minX, v.X)
		minY = math.Min(minY, v.Y)
	}

	// 在范围内随机初始化k个中心点
	for i := 0; i < k.k; i++ {
		x := rand.Float64()*(maxX-minX) + minX
		y := rand.Float64()*(maxY-minY) + minY
		k.clusters[i].Centroide = Point{X: x, Y: y}
	}
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
