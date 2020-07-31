package point

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestFromRequest(t *testing.T) {
	tables := []struct {
		params, res string
	}{
		{"", "x value not provided"},
		{"?x=3", "y value not provided"},
		{"?x=3&y=4", "(3,4)"},
	}

	for _, table := range tables {
		url := fmt.Sprintf("/api/points%s", table.params)
		p, err := FromRequest(httptest.NewRequest("GET", url, nil))
		if err != nil {
			if s := err.Error(); s != table.res {
				t.Errorf("Resulting error of %s was incorrect, got: %s, want: %s", url, s, table.res)
			}
			continue
		}

		if s := p.String(); s != table.res {
			t.Errorf("Result of %s was incorrect, got: %s, want: %s", url, s, table.res)
		}
	}
}
