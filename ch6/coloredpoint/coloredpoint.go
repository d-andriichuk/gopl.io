package coloredpoint

import (
	"image/color"

	"gopl.io/ch6/geometry"
)

//
// Struct Colored point
//
type ClolredPoint struct {
	geometry.Point
	Color color.RGBA
}
