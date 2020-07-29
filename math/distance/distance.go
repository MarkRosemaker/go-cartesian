package distance

import (
	"net/http"

	"github.com/MarkRosemaker/go-server/server/form"
)

// FromRequest returns the from the form value or an error, if a distance couldn't be parsed.
func FromRequest(req *http.Request) (int, error) {
	return form.GetIntE(req, "distance")
}
