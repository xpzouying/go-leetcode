package two_sum

import "testing"

func TestTwoSum(t *testing.T) {

	ts := []struct {
		List   []int
		Target int
		Want   []int
	}{
		{
			List:   []int{2, 7, 11, 15},
			Target: 9,
			Want:   []int{0, 1},
		},
		{
			List:   []int{},
			Target: 9,
			Want:   []int{},
		},
		{
			List:   []int{1},
			Target: 9,
			Want:   []int{},
		},
		{
			List:   []int{1, 2},
			Target: 3,
			Want:   []int{0, 1},
		},
		{
			List:   []int{1, 2},
			Target: 4,
			Want:   []int{},
		},

		// 无序的情况
		{
			List:   []int{3, 2, 4},
			Target: 6,
			Want:   []int{1, 2},
		},
	}

	for _, tc := range ts {

		got := twoSum(tc.List, tc.Target)

		if !equalList(tc.Want, got) {
			t.Errorf("want: %v, got: %v\n", tc.Want, got)
		}
	}
}

func equalList(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
