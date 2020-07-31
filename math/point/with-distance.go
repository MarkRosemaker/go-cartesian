package point

// WithDistance is a container for a point and its distance to a search origin point.
type WithDistance struct {
	Point    `json:"point"`
	Distance int `json:"distance"`
}
