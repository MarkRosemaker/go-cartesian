// Package distance offers functions to calculate the distance between two points.
//
// For now, there's only one function.
package distance

import (
	"github.com/MarkRosemaker/go-cartesian/math/point"
)

// Manhattan calculates the Manhatten distance between two points.
func Manhattan(p1 point.Point, p2 point.Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

// abs calculates the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
