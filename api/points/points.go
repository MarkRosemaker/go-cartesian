// Package points defines the behavior of the /api/points API.
package points

import (
	"net/http"

	"github.com/MarkRosemaker/go-cartesian/math/point"
	"github.com/MarkRosemaker/go-cartesian/math/points"
	"github.com/MarkRosemaker/go-cartesian/math/radius"

	"github.com/MarkRosemaker/go-server/server/api"
	"github.com/MarkRosemaker/go-server/server/context"
)

// Respond responds to API requests to /api/points.
//
// The returning interface will be marshalled into a JSON.
// If there is an error, an instance of type api.Error will get returned.
//
// In addition to the challenge requirements, this function will timeout
// if it takes longer than a timeout that was given as a parameter.
func Respond(req *http.Request) interface{} {
	ctx, cancel := context.WithUserTimeout(req)
	defer cancel()

	origin, err := point.FromRequest(req)
	if err != nil {
		return api.ErrBadRequest(err)
	}

	var r int
	if r, err = radius.FromRequest(req); err != nil {
		return api.ErrBadRequest(err)
	}

	// start finding neighbors
	var nb points.Points
	success := make(chan bool)
	go func() <-chan bool {
		nb = pts.Neighbors(*origin, r)
		success <- true
		return success
	}()

	// timout if necessary
	select {
	case <-success:
		return nb
	case <-ctx.Done():
		return api.ErrWrap(ctx.Err())
	}
}
