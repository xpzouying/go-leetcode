package min_cost_climbing_stairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	ts := []struct {
		Cost []int
		Want int
	}{
		{
			Cost: []int{10, 15, 20},
			Want: 15,
		},
		{
			Cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			Want: 6,
		},
	}

	for _, tc := range ts {
		got := minCostClimbingStairs(tc.Cost)
		if tc.Want != got {
			t.Errorf("cost list: %v, want=%d, got=%d",
				tc.Cost, tc.Want, got)
		}
	}
}
