package triangle

import "testing"

func TestTriangle(t *testing.T) {
	ts := []struct {
		Triangle [][]int
		Want     int
	}{
		{
			Triangle: [][]int{
				[]int{2},
				[]int{3, 4},
				[]int{6, 5, 7},
				[]int{4, 1, 8, 3},
			},
			Want: 11,
		},
		{
			Triangle: [][]int{
				[]int{1},
				[]int{2, 3},
			},
			Want: 3,
		},
	}

	for _, tc := range ts {
		got := minimumTotal(tc.Triangle)
		if tc.Want != got {
			t.Errorf("input: %v, want: %d, got: %d", tc.Triangle, tc.Want, got)
		}
	}

}
