package points

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// move up to repo level
	os.Chdir("../..")

	err := Init(false)
	if err != nil {
		t.Fatal(err)
	}

	if len(pts) == 0 {
		t.Error("parsed points.json file but didn't find any points")
	}
}
