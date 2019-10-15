package move_zeroes

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {

	ts := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{0, 1, 0, 3, 12},
			want:  []int{1, 3, 12, 0, 0},
		},
		{
			input: []int{1, 2, 3, 0, 0, 0},
			want:  []int{1, 2, 3, 0, 0, 0},
		},
		{
			input: []int{1, 0, 2, 0, 3},
			want:  []int{1, 2, 3, 0, 0},
		},
		{
			input: []int{0, 1, 0, 2},
			want:  []int{1, 2, 0, 0},
		},
	}

	for _, tc := range ts {
		moveZeroes(tc.input)
		if sliceIsEqual(tc.want, tc.input) == false {
			t.Errorf("want:%+v, got:%+v", tc.want, tc.input)
		}
	}
}

func sliceIsEqual(a, b []int) bool {
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
