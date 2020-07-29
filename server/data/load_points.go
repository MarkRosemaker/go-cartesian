package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/MarkRosemaker/go-cartesian/points"
)

// the file we are importing from
const pointsFile string = `data/points.json`

// LoadPoints reads a list of points from a (pre-defined) file and returns it.
func loadPoints() (points.Points, error) {

	// read the json file into bytes
	b, err := ioutil.ReadFile(pointsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("couldn't find points file %q, please provide one", pointsFile)
		}
		return nil, fmt.Errorf("problem reading points file %q: %s", pointsFile, err)
	}

	// initialize points
	var pts points.Points

	// unmarshal json
	if err = json.Unmarshal(b, &pts); err != nil {
		return nil, fmt.Errorf("problem unmarshalling json file %q: %s", pointsFile, err)
	}

	return pts, nil
}
