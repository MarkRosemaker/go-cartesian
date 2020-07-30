package points

import (
	"github.com/MarkRosemaker/go-cartesian/math/distance"
	"github.com/MarkRosemaker/go-cartesian/math/point"
)

// Points is an array of points.
type Points []point.Point

// Neighbors returns a slice of all points within radius of an origin point.
// They are sorted in order of increasing distance.
func (pts Points) Neighbors(origin point.Point, radius int) Points {
	return pts.NeighborsWithDistance(origin, radius).Points()
}

// NeighborsWithDistance returns a list of points within radius of this point, along with the distance of those points to the origin point.
// They are sorted in order of increasing distance.
func (pts Points) NeighborsWithDistance(origin point.Point, radius int) WithDistance {
	// initialize a list of points along with their distance to this point
	k := len(pts)
	list := make(WithDistance, k)

	// make that list
	for i := 0; i < k; i++ {
		list[i] = point.WithDistance{
			Point:    pts[i],
			Distance: distance.Between(pts[i], origin),
		}
	}

	// "The points should be returned in order of increasing distance from the search origin."

	// todo: which is faster for long lists of points, sortCut or collectSort?
	// it might depend on the size of the list
	return list.sortCut(radius)
	// return list.collectSort(radius)
}
