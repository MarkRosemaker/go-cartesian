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
