package rotate_array

import "testing"

func TestRotate(t *testing.T) {

	ts := []struct {
		input []int
		k     int
		want  []int
	}{

		{
			input: []int{1, 2, 3, 4, 5, 6, 7},
			k:     3,
			want:  []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			input: []int{-1, -100, 3, 99},
			k:     2,
			want:  []int{3, 99, -1, -100},
		},
		{
			input: []int{1, 2, 3, 4, 5, 6, 7},
			k:     3,
			want:  []int{5, 6, 7, 1, 2, 3, 4},
		},
	}

	for _, tc := range ts {
		orig := make([]int, len(tc.input))
		copy(orig, tc.input)

		rotate(tc.input, tc.k)
		if sliceIsEqual(tc.want, tc.input) == false {
			t.Errorf("input:%v, k:%v, want:%v, got:%v",
				orig, tc.k, tc.want, tc.input)
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
