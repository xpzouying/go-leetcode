package perfect_squares

import "testing"

func TestNumSquares(t *testing.T) {

	ts := []struct {
		input int
		want  int
	}{
		{input: 12, want: 3}, // 4+4+4
		{input: 13, want: 2}, // 4+9
		{input: 4, want: 1},  // 4
		{input: 5, want: 2},  // 4+1
		{input: 9, want: 1},  // 9
		{input: 8, want: 2},  // 4+4
	}

	for _, tc := range ts {
		got := numSquares(tc.input)
		if tc.want != got {
			t.Errorf("input:%d, want:%d, got:%d", tc.input, tc.want, got)
		}
	}

}
