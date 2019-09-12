package best_time_to_buy_and_sell_stock_ii

import "testing"

func TestMaxProfit(t *testing.T) {

	ts := []struct {
		input []int
		want  int
	}{
		{
			input: []int{7, 1, 5, 3, 6, 4},
			want:  7,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  4,
		},
		{
			input: []int{1, 2, 3, 1, 5},
			want:  6,
		},
		{
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
		{
			input: []int{1, 3},
			want:  2,
		},
		{
			input: []int{3, 1, 2},
			want:  1,
		},
	}

	for _, tc := range ts {
		got := maxProfit(tc.input)
		if tc.want != got {
			t.Errorf("input: %v, want: %d, got: %d",
				tc.input, tc.want, got)
		}
	}

}
