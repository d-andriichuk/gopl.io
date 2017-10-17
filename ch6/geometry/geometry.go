package geometry

import (
	"math"
)

//
// Type Path
//
type Path []Point

//
// Type Point
//
type Point struct{ X, Y float64 }

//
// Distance returns way length
//
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//
// Traditional function
//
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//
// Method Distance
//
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//
// Method ScaleBy
//
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
