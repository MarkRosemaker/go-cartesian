package serverdata

import (
	"fmt"

	"github.com/MarkRosemaker/go-cartesian/point"
)

// data is our singleton data object
var data Data = Data{}

// Get returns a pointer to the singleton data object.
func Get() *Data {
	return &data
}

// Data represents data that lies on the server.
type Data struct {
	AllPoints point.Points
}

// Load loads the server data to a singleton object.
func Load() error {
	pts, err := loadPoints()
	if err != nil {
		return fmt.Errorf("error loading points: %s", err)
	}
	data.AllPoints = pts

	// later we can load other stuff here

	return nil
}
