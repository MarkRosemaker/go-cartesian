package radius

import (
	"net/http"

	"github.com/MarkRosemaker/go-server/server/form"
)

// FromRequest returns the radius (or 'distance') from the form value or an error, if  'distance' couldn't be parsed.
func FromRequest(req *http.Request) (int, error) {
	return form.GetIntE(req, "distance")
}
