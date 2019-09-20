package coin_change

import "testing"

func TestCoinChange(t *testing.T) {

	ts := []struct {
		coins  []int
		amount int
		want   int
	}{
		{
			coins:  []int{1, 2, 5},
			amount: 4,
			want:   2,
		},
		{
			coins:  []int{1, 2},
			amount: 3,
			want:   2,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			coins:  []int{6, 2, 5},
			amount: 11,
			want:   2,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 1,
			want:   1,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 2,
			want:   1,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 5,
			want:   1,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 3,
			want:   2,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 6,
			want:   2,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 7,
			want:   2,
		},
		{
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
	}

	for _, tc := range ts {
		got := coinChange(tc.coins, tc.amount)
		if tc.want != got {
			t.Errorf("coins: %v, amount:%d want:%d got:%d",
				tc.coins, tc.amount, tc.want, got)
		}
	}
}
