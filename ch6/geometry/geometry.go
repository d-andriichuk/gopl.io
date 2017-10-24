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
// Add method
//
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

//
// Sub method
//
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

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
// TranslateBy method
//
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
