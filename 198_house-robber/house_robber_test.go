package house_robber

import "testing"

func TestRob(t *testing.T) {
	ts := []struct {
		Input []int
		Want  int
	}{

		{
			Input: []int{1, 2, 3, 1},
			Want:  4,
		},
		{

			Input: []int{2, 7, 9, 3, 1},
			Want:  12,
		},
		{
			Input: []int{},
			Want:  0,
		},
		{
			Input: []int{1},
			Want:  1,
		},
		{
			Input: []int{1, 2},
			Want:  2,
		},
		{
			Input: []int{2, 1, 1, 2},
			Want:  4,
		},
	}

	for _, tc := range ts {
		got := rob(tc.Input)
		if tc.Want != got {
			t.Errorf("rob: %+v, want: %d, got: %d", tc.Input, tc.Want, got)
		}

	}

}
