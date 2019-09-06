package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {

	ts := []struct {
		Stairs int
		Want   int
	}{

		{
			Stairs: 1,
			Want:   1,
		},
		{
			Stairs: 2,
			Want:   2,
		},
		{
			Stairs: 3,
			Want:   3,
		},
	}

	for _, tc := range ts {
		got := climbStairs(tc.Stairs)
		if tc.Want != got {
			t.Errorf("climb stairs - %d, expected: %d, got: %d", tc.Stairs, tc.Want, got)
		}

	}

}
