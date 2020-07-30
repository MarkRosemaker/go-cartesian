package point

// WithDistance is a container for a point and the distance to a search origin point.
type WithDistance struct {
	Point    `json:"point"`
	Distance int `json:"distance"`
}
