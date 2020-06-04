package count_squares

import (
	"testing"
)

func TestCountSquares(t *testing.T) {

	ts := []struct {
		matrix [][]int
		want   int
	}{
		{

			[][]int{
				[]int{0, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{0, 1, 1, 1},
			},
			15,
		},
		{
			[][]int{
				[]int{1, 0, 1},
				[]int{1, 1, 0},
				[]int{1, 1, 0},
			},
			7,
		},
	}

	for _, tc := range ts {
		matrix := tc.matrix
		got := countSquares(matrix)
		if tc.want != got {
			t.Errorf("matrix=%+v, want=%d, got=%d", matrix, tc.want, got)
		}
	}
}
