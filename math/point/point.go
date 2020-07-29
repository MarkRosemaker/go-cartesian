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

func FromRequest(req *http.Request) (*Point, error) {
	p := &Point{}

	if p.X, err := form.GetIntE(req, "x"); err != nil {
		return nil, err
	}

	if p.Y, err := form.GetIntE(req, "y"); err != nil {
		return nil, err
	}

	return p, nil
}
