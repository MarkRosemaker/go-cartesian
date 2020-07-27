// Package data represents the data *that is given to a template*.
package pagedata

import (
	"net/http"
	"strconv"

	"github.com/MarkRosemaker/go-cartesian/point"
	"github.com/MarkRosemaker/go-cartesian/server/serverdata"
)

// Data is the data object that will be given to the template of a page.
type Data struct {
	Req *http.Request
	// Error error
	*serverdata.Data
}

// New creates a new data object from a request.
func New(req *http.Request) *Data {
	return &Data{req, serverdata.Get()}
}

// PointsWithinDistance returns all points within distance.
func (d *Data) PointsWithinDistance() point.Points {
	p, err := point.FromRequest(d.Req)
	if err != nil {
		return make(point.Points, 0)
	}

	var dist int
	if dist, err = d.searchDistanceWithError(); err != nil {
		return make(point.Points, 0)
	}

	return p.WithinDistance(d.AllPoints, dist)
}

// SearchOriginPoint returns the point that the user is searching for.
// It can be used in templates by using 'with' as the error is discarded.
func (d *Data) SearchOriginPoint() *point.Point {
	p, _ := point.FromRequest(d.Req)
	return p
}

// SearchDistance returns the distance that the user is searching for.
// It can be used in templates by using 'with' as the error is discarded.
func (d *Data) SearchDistance() int {
	dist, _ := d.searchDistanceWithError()
	return dist
}

// searchDistanceWithError returns the distance that the user is searching for, or an error.
func (d *Data) searchDistanceWithError() (int, error) {
	dist, err := strconv.Atoi(d.Req.FormValue("distance"))
	if err != nil {
		return 0, err
	}
	return dist, nil
}
