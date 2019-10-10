package number_of_1_bits

import "testing"

func TestHammingWeight(t *testing.T) {
	ts := []struct {
		input uint32
		want  int
	}{
		{input: 11, want: 3},
		{input: 1, want: 1},
		{input: 2, want: 1},
		{input: 0, want: 0},
		{input: 3, want: 2},
	}

	for _, tc := range ts {
		got := hammingWeight(tc.input)
		if tc.want != got {
			t.Errorf("input:%d, want:%d, got:%d", tc.input, tc.want, got)
		}
	}

}
