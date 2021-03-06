// Package point defines behavior related to single point objects.
package point

import (
	"fmt"
	"net/http"

	"github.com/MarkRosemaker/go-server/server/form"
)

// Point represents a point on a cartesian plane.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// String is defined for testing, display and debugging purposes.
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// FromRequest returns a point by parsing the form values "x" and "y", or an error, if they couldn't be parsed.
func FromRequest(req *http.Request) (*Point, error) {

	var err error
	p := &Point{}

	if p.X, err = form.GetIntE(req, "x"); err != nil {
		return nil, err
	}

	if p.Y, err = form.GetIntE(req, "y"); err != nil {
		return nil, err
	}

	return p, nil
}
