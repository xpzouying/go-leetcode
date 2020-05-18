package three_steps_problem

import (
	"testing"
)

func TestWaysToStep(t *testing.T) {

	ts := []struct {
		n    int
		want int
	}{
		{3, 4},
		{5, 13},
		{10, 274},
		{50, 230552708},
	}

	for _, tc := range ts {
		got := waysToStep(tc.n)
		if tc.want != got {
			t.Errorf("n=%d, want=%d, but got=%d", tc.n, tc.want, got)
		}
	}
}
