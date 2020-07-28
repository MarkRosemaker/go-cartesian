package point

import (
	"fmt"
)

// Point represents a point on a cartesian plane.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// String is defined for display and debugging purposes.
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// // DistanceTo calculates the Manhattan distance between this point and another.
// func (p Point) DistanceTo(p2 Point) int {
// 	return abs(p.X-p2.X) + abs(p.Y-p2.Y)
// }

// // abs calculates the absolute value of an integer
// func abs(x int) int {
// 	if x < 0 {
// 		return -1 * x
// 	}
// 	return x
// }

// // Neighbors returns a slice of all points within radius of this point.
// // They are sorted in order of increasing distance.
// func (p Point) Neighbors(pts Points, radius int) Points {
// 	return p.NeighborsWithDistance(pts, radius).Points()
// }

// // NeighborsWithDistance returns a list of points within radius of this point, along with the distance of those points to the origin point.
// // They are sorted in order of increasing distance.
// func (p Point) NeighborsWithDistance(pts Points, radius int) PointsWithDistance {
// 	// initialize a list of points along with their distance to this point
// 	k := len(pts)
// 	l := make(PointsWithDistance, k)

// 	for i := 0; i < k; i++ {
// 		l[i] = WithDistance{pts[i], p.Distance(*pts[i])}
// 	}

// 	log.Println("neighbors with distance:", l)

// 	// "The points should be returned in order of increasing distance from the search origin."

// 	// sort the list
// 	sort.SliceStable(l, func(i int, j int) bool {
// 		return l[i].Distance < l[j].Distance
// 	})

// 	// cut off all that are too far away
// 	for i := 0; i < k; i++ {
// 		if l[i].Distance > radius {
// 			return l[:i]
// 		}
// 	}

// 	return l
// }
