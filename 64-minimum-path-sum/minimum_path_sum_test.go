package minimum_path_sum

import "testing"

func TestMinPathSum(t *testing.T) {

	ts := []struct {
		M    [][]int
		Want int
	}{
		{
			M: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			Want: 7,
		},
		{
			M: [][]int{
				[]int{1, 2},
			},
			Want: 3,
		},
		{
			M: [][]int{
				[]int{1},
				[]int{2},
			},
			Want: 3,
		},
	}

	for _, tc := range ts {
		got := minPathSum(tc.M)
		if tc.Want != got {
			t.Errorf("matrix: %+v, want: %d, got: %d",
				tc.M, tc.Want, got)
		}
	}
}
