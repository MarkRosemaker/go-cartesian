package points

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/MarkRosemaker/go-cartesian/math/points"
)

// the file we are importing from
const pointsFile string = `data/points.json`

// the points that we imported
var pts points.Points

// Init reads a list of points from 'data/points.json' into memory.
func Init(verbose bool) error {

	// read the json file into bytes
	b, err := ioutil.ReadFile(pointsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("couldn't find points file %q, please provide one", pointsFile)
		}
		return fmt.Errorf("problem reading points file %q: %s", pointsFile, err)
	}

	// unmarshal json
	if err = json.Unmarshal(b, &pts); err != nil {
		return fmt.Errorf("problem unmarshalling json file %q: %s", pointsFile, err)
	}

	if verbose {
		log.Printf("loaded %d points", len(pts))
	}

	return nil
}
