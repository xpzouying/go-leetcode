package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {

	ts := []struct {
		Nums []int
		Val  int
		Want int
	}{

		{
			Nums: []int{3, 2, 2, 3},
			Val:  3,
			Want: 2,
		},
		{
			Nums: []int{1},
			Val:  3,
			Want: 1,
		},
		{
			Nums: []int{1},
			Val:  1,
			Want: 0,
		},
	}

	for _, tc := range ts {
		got := removeElement(tc.Nums, tc.Val)
		if got != tc.Want {
			t.Errorf("nums=%v, Val=%d, Want=%d, Got=%d",
				tc.Nums, tc.Val, tc.Want, got)
		}
	}

}
