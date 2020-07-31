// Package radius contains a method to get the search radius from user input.
//
// The point of this package is just to have standardized and beautiful code. :)
package radius

import (
	"errors"
	"net/http"

	"github.com/MarkRosemaker/go-server/server/form"
)

// FromRequest returns the radius (or "distance") from the form value or an error, if "distance" couldn't be parsed or is negative.
func FromRequest(req *http.Request) (int, error) {
	r, err := form.GetIntE(req, "distance")
	if err != nil {
		return r, err
	}

	if r < 0 {
		return 0, errors.New("negative distance")
	}
	return r, err
}
