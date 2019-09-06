package maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	ts := []struct {
		List []int
		Sum  int
	}{
		{
			List: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			Sum:  6,
		},
		{
			List: []int{1},
			Sum:  1,
		},
		{
			List: []int{3, -1, 2},
			Sum:  4,
		},
	}

	for _, tc := range ts {
		got := maxSubArray(tc.List)
		if got != tc.Sum {
			t.Errorf("got max error: list=%+v, want=%d, got=%d",
				tc.List, tc.Sum, got)
		}
	}

}
