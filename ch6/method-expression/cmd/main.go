package main

import (
	"fmt"

	"gopl.io/ch6/geometry"
)

func main() {
	var path = geometry.Path{
		geometry.Point{1, 1},
		geometry.Point{2, 3},
		geometry.Point{4, 7},
	}

	var offset = geometry.Point{2, 2}

	path.TranslateBy(offset, true)

	fmt.Println(path)
}
