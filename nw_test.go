package align_test

import (
	"testing"

	"github.com/gregoryv/align"
)

func TestNeedlemanWunsch(t *testing.T) {
	cases := []struct {
		a, b []rune
	}{
		{[]rune("abc"), []rune("123")},
	}
	for i, c := range cases {
		result := align.NeedlemanWunsch(c.a, c.b)
		if result == nil {
			t.Errorf("%v. expected a result", i)
		}
	}
}
