package single_number

import "testing"

func TestSingleNumber(t *testing.T) {

	ts := []struct {
		Nums []int
		Want int
	}{
		{
			Nums: []int{2, 2, 1},
			Want: 1,
		},
		{
			Nums: []int{4, 1, 2, 1, 2},
			Want: 4,
		},
	}

	for _, tc := range ts {
		got := singleNumber(tc.Nums)
		if got != tc.Want {
			t.Errorf("single number: nums=%v, want=%d, got=%d", tc.Nums, tc.Want, got)
		}
	}
}
