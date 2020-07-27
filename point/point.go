package point

// Point represent a point on a cartesian plane.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Points is an array of points.
type Points []Point

// Dist calculates the Manhattan distance between this point and another.
func (p Point) Dist(p2 Point) int {
	return abs(p.X-p2.X) + abs(p.Y-p2.Y)
}

// abs calculates the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// WithinDistance returns all the points within distance of this point.
func (p Point) WithinDistance(pts Points, dist int) PointsWithDistance {
	// initialize a list of points along with their distance to this point
	k := len(pts)
	l := make(PointsWithDistance, k)

	for i := 0; i < k; i++ {
		l[i] = WithDistance{&pts[i], p.Dist(pts[i])}
	}

	// The points should be returned in order of increasing distance from the search origin.
	return l.sortAndCut(dist)
}
