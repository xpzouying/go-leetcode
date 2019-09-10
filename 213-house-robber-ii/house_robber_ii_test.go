package house_robber_ii

import "testing"

func TestHouseRobber(t *testing.T) {
	ts := []struct {
		Input []int
		Want  int
	}{

		{
			Input: []int{2, 3, 2},
			Want:  3,
		},
		{
			Input: []int{1, 2, 3, 1},
			Want:  4,
		},
		{
			Input: []int{1, 2, 3, 4},
			Want:  6,
		},
		{
			Input: []int{2, 3, 4},
			Want:  4,
		},
		{
			Input: []int{1, 3, 1, 3, 100},
			Want:  103,
		},
	}

	for _, tc := range ts {

		got := rob(tc.Input)
		if tc.Want != got {
			t.Errorf("input: %+v, want=%d, got=%d", tc.Input, tc.Want, got)
		}
	}

}
