package point

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	notProvided = "%s value not provided"
	notParsed   = "%s value %q could not be parsed to int"
)

// FromRequest returns a Point from the form values or an error, if a point couldn't be parsed.
func FromRequest(req *http.Request) (*Point, error) {

	var (
		xStr, yStr string
		x, y       int
		err        error
	)

	if xStr := req.FormValue("x"); xStr == "" {
		return nil, fmt.Errorf(notProvided, "x")
	}
	if yStr := req.FormValue("y"); yStr == "" {
		return nil, fmt.Errorf(notProvided, "y")
	}
	if x, err = strconv.Atoi(xStr); err != nil {
		return nil, fmt.Errorf(notParsed, "x", xStr)
	}
	if y, err = strconv.Atoi(yStr); err != nil {
		return nil, fmt.Errorf(notParsed, "y", yStr)
	}

	return &Point{x, y}, nil
}
