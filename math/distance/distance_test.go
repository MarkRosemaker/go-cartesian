package distance

import (
	"testing"

	"github.com/MarkRosemaker/go-cartesian/math/point"
)

func TestBetween(t *testing.T) {
	tables := []struct {
		x1, y1, x2, y2, dist int
	}{
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 2},
		{1, -1, 0, 0, 2},
		{-5, -5, 0, 5, 15},
		{1, 6, -1, 5, 3},
		{3, 5, -1, 5, 4},
		{2, 3, -1, 5, 5},
		{2, 3, 3, 5, 3},
		{3, 4, 1, 6, 4},
		{2, 3, 1, 6, 4},
		{-123456789, 0, 0, 0, 123456789},
		{1000000000000000000, 1000000000000000000, -1000000000000000000, -1000000000000000000, 4000000000000000000},
	}

	for _, table := range tables {
		a := point.Point{X: table.x1, Y: table.y1}
		b := point.Point{X: table.x2, Y: table.y2}
		if res := Between(a, b); res != table.dist {
			t.Errorf("Result of %s, %s was incorrect, got: %d, want: %d.", a, b, table.dist, res)
		}
	}
}
