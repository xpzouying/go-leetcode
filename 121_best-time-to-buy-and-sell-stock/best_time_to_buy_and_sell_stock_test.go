package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {

	ts := []struct {
		Input []int
		Want  int
	}{

		{
			Input: []int{7, 1, 5, 3, 6, 4},
			Want:  5,
		},
		{
			Input: []int{7, 1},
			Want:  0,
		},
		{
			Input: []int{1, 7},
			Want:  6,
		},
		{
			Input: []int{1, 7, 3, 8},
			Want:  7,
		},
	}

	for _, tc := range ts {

		got := maxProfit(tc.Input)
		if tc.Want != got {
			t.Errorf("input: %+v, want: %d, got: %d", tc.Input, tc.Want, got)
		}
	}
}
