package maxvalue

import "testing"

func TestMaxValue(t *testing.T) {

	ts := []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 12,
		},
	}

	for _, tc := range ts {
		got := maxValue(tc.input)
		if tc.want != got {
			t.Fatalf("input: %+v, want: %d, got: %d", tc.input, tc.want, got)
		}
	}
}
