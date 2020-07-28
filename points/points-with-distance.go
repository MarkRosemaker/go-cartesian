package points

// WithDistance is a container for a point and the distance to the search origin point.
// type WithDistance struct {
// 	*Point   `json:"point"`
// 	Distance int `json:"distance"`
// }

// // PointsWithDistance is a slice of points with their search origin point.
// type PointsWithDistance []WithDistance

// // Points returns just the points, without the distance to the origin point.
// func (pwd PointsWithDistance) Points() Points {
// 	k := len(pwd)

// 	o := make(Points, k)
// 	for i := 0; i < k; i++ {
// 		o[i] = pwd[i].Point
// 	}
// 	return o
// }
