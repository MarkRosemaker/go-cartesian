// Package radius contains a method to get the search radius from user input.
//
// The point of this package is just to have standardized and beautiful code. :)
package radius

import (
	"net/http"

	"github.com/MarkRosemaker/go-server/server/form"
)

// FromRequest returns the radius (or distance) from the form value or an error, if 'radius' and 'distance' couldn't be parsed.
func FromRequest(req *http.Request) (int, error) {
	if r, err := form.GetIntE(req, "radius"); err == nil {
		return r, err
	}
	return form.GetIntE(req, "distance")
}
