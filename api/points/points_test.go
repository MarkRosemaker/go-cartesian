// Package points defines the behavior of the /api/points API.
package points

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MarkRosemaker/go-cartesian/math/points"
	"github.com/MarkRosemaker/go-server/server/api"
)

func TestRespond(t *testing.T) {
	os.Chdir("../..")
	Init(false)

	tables := []struct {
		params string
		res    string
	}{
		// test all errors
		{"",
			"400 Bad Request: x value not provided"},
		{"?x=foo",
			"400 Bad Request: x value 'foo' could not be parsed to int"},
		{"?x=100000000000000000000000000000000000000000",
			"400 Bad Request: x value '100000000000000000000000000000000000000000' could not be parsed to int"},
		{"?x=3",
			"400 Bad Request: y value not provided"},
		{"?x=3&y=blub",
			"400 Bad Request: y value 'blub' could not be parsed to int"},
		{"?x=3&y=4",
			"400 Bad Request: distance value not provided"},
		{"?x=3&y=4&distance=10&timeout=0s",
			"408 Request Timeout"},
		{"?x=0&y=0&distance=-1",
			"400 Bad Request: negative distance"},

		{"?x=0&y=0&distance=0", "[]"},
		{"?x=0&y=0&distance=9", "[]"},
		{"?x=0&y=0&distance=10", "[(2,-8)]"},
		{"?x=-100&y=-100&distance=40", "[(-92,-93),(-88,-72)]"},

		// all points, note that this test might fail if two points with the same distance are swapped
		// a more sophisticated testing method is needed
		{"?x=3&y=4&distance=400",
			"[(2,-8),(18,-3),(-6,-10),(-20,2),(27,1),(-9,-12),(-25,11),(-3,-29),(17,37),(19,35),(-13,-29),(30,26),(9,49),(51,8),(7,-44),(8,52),(-3,52),(14,50),(58,-8),(-41,29),(16,62),(37,-33),(-7,-59),(64,17),(-30,-38),(-60,16),(-36,41),(29,55),(78,8),(-69,-4),(-55,29),(-25,60),(-3,-76),(-24,63),(54,-33),(-46,45),(-13,-72),(-31,63),(-84,-2),(43,-50),(-59,38),(47,-51),(-5,95),(-38,-59),(83,-20),(-64,-33),(74,37),(-65,-37),(95,-14),(-32,79),(79,39),(76,42),(-17,96),(-46,67),(-61,-45),(-26,-80),(-43,-65),(36,86),(-44,-65),(28,95),(-47,71),(39,-78),(-44,75),(84,-33),(-50,-61),(-80,41),(62,-57),(-61,60),(-41,-75),(-32,-84),(92,-32),(90,43),(-77,-45),(74,65),(-73,60),(91,49),(-60,74),(-29,-99),(63,-72),(-82,55),(-42,97),(-42,-92),(83,66),(54,-94),(96,64),(73,89),(94,-60),(84,-71),(-89,69),(-74,-77),(88,-74),(80,-83),(-88,-72),(-77,93),(-98,74),(-89,86),(-94,89),(-100,84),(98,-86),(-92,-93)]"},
	}

	for _, table := range tables {
		url := fmt.Sprintf("/api/points%s", table.params)
		resp := Respond(httptest.NewRequest("GET", url, nil))

		switch v := resp.(type) {
		case api.Error:
			if s := v.Error(); s != table.res {
				t.Errorf("Result of %s was incorrect, got: %s, want: %s.",
					url, s, table.res)
			}
		case points.Points:
			if s := fmt.Sprintf("%v", v); s != table.res {
				t.Errorf("Result of %s was incorrect, got: %s, want: %s", url, s, table.res)
			}
		default:
			t.Errorf("Result of %s has wrong type, expected: api.Error or points.Points, got: %T", url, resp)
		}
	}
}
