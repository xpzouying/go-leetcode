package longest_increasing_subsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {

	ts := []struct {
		input []int
		want  int
	}{
		{
			input: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want:  4,
		},
		{
			input: []int{3, 2, 4, 7, 5, 8, 6},
			want:  4,
		},
		{
			input: []int{3, 2},
			want:  1,
		},
		{
			input: []int{3, 2, 4, 1},
			want:  2,
		},
		{
			input: []int{4, 10, 4, 3, 8, 9},
			want:  3,
		},
	}

	for _, tc := range ts {
		got := lengthOfLIS(tc.input)
		if tc.want != got {
			t.Errorf("input: %+v, want: %d, got: %d",
				tc.input, tc.want, got)
		}
	}

}
