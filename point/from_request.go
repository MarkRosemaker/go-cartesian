package point

import (
	"errors"
	"net/http"
	"strconv"
)

// FromRequest returns a Point from the form values or an error, if a point couldn't be parsed.
func FromRequest(req *http.Request) (*Point, error) {

	xStr := req.FormValue("x")
	if xStr == "" {
		return nil, errors.New("please provide an x value")
	}

	yStr := req.FormValue("y")
	if yStr == "" {
		return nil, errors.New("please provide a y value")
	}

	var (
		x, y int
		err  error
	)

	if x, err = strconv.Atoi(xStr); err != nil {
		return nil, err
	}
	if y, err = strconv.Atoi(yStr); err != nil {
		return nil, err
	}

	return &Point{x, y}, nil
}
