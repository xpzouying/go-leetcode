package plus_one

import (
	"testing"
)


func TestPlusOne(t *testing.T) {
	ts := []struct {
		input []int
		want []int
	} {
		{
			input: []int{1,2,3},
			want: []int{1,2,4},
		},
		{
			input: []int{1},
			want: []int{2},
		},
		{
			input: []int{0},
			want: []int{1},
		},
		{
			input: []int{9},
			want: []int{1, 0},
		},
		{
			input: []int{9, 9},
			want: []int{1, 0, 0},
		},
	}

	for _, tc := range ts {
		got := plusOne(tc.input)
		if sliceIsEqual(tc.want, got) == false {
			t.Errorf("plus one error: input:%v, want:%v, got:%v",
				tc.input, tc.want, got)
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
