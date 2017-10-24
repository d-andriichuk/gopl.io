package main

import (
	"fmt"
	"image/color"

	"gopl.io/ch6/coloredpoint"
	"gopl.io/ch6/geometry"
)

func main() {
	var cp coloredpoint.ClolredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)

	cp.Y = 2
	fmt.Println(cp.Point.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = coloredpoint.ClolredPoint{geometry.Point{1, 1}, red}
	var q = coloredpoint.ClolredPoint{geometry.Point{5, 4}, blue}

	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}
