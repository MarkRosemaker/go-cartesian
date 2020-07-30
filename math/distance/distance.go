package distance

import (
	"github.com/MarkRosemaker/go-cartesian/math/point"
)

// Between calculates the Manhatten distance between two points.
func Between(p1 point.Point, p2 point.Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

// abs calculates the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
