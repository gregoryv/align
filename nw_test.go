package align_test

import (
	"github.com/gregoryv/align"
	"testing"
)

func TestNeedlemanWunsch(t *testing.T) {
	empty := make([]rune, 0)
	cases := []struct {
		a, b []rune
	}{
		{empty, empty},
		{[]rune("abc"), empty},
	}
	for i, c := range cases {
		result := align.NeedlemanWunsch(c.a, c.b)
		if result == nil {
			t.Errorf("%v. expected a result got nil", i)
		}
	}

}
