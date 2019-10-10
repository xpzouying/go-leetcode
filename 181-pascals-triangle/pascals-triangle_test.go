package pascals_triangle

import "testing"

func TestGenerate(t *testing.T) {

	ts := []struct {
		input int
		want  [][]int
	}{
		{
			input: 1,
			want: [][]int{
				[]int{1},
			},
		},
		{
			input: 2,
			want: [][]int{
				[]int{1},
				[]int{1, 1},
			},
		},
		{
			input: 3,
			want: [][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
			},
		},
		{
			input: 4,
			want: [][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
			},
		},
		{
			input: 5,
			want: [][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
	}

	for _, tc := range ts {
		got := generate(tc.input)
		if intSliceIsEqual(tc.want, got) == false {
			t.Errorf("input:%d, want:%+v, got:%+v", tc.input, tc.want, got)
		}
	}
}

func intSliceIsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {

		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
