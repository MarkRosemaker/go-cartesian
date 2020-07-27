package point

import "sort"

// WithDistance is a container for a point and the distance to the search origin point.
type WithDistance struct {
	*Point
	dist int
}

// PointsWithDistance is a slice of points with their search origin point.
type PointsWithDistance []WithDistance

// Satisfy the sort interface

// Len is the number of elements in the collection.
func (pts PointsWithDistance) Len() int {
	return len(pts)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pts PointsWithDistance) Less(i, j int) bool {
	return pts[i].dist < pts[j].dist
}

// Swap swaps the elements with indexes i and j.
func (pts *PointsWithDistance) Swap(i, j int) {
	pts[i], pts[j] = pts[j], pts[i]
}

func (pts *PointsWithDistance) sortAndCut(dist int) PointsWithDistance {
	sort.Sort(pts)
	return pts
}
