package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {

	ts := []struct {
		Nums []int
		Want bool
	}{
		{
			Nums: []int{1, 2, 3, 1},
			Want: true,
		},
		{
			Nums: []int{1, 2, 3, 4},
			Want: false,
		},
		{
			Nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			Want: true,
		},
	}

	for _, tc := range ts {
		got := containsDuplicate(tc.Nums)
		if got != tc.Want {
			t.Errorf("nums: %v, want: %v, got: %v",
				tc.Nums, tc.Want, got)
		}
	}

}
