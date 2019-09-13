package power_of_two

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	ts := []struct {
		input int
		want  bool
	}{
		{
			input: 1,
			want:  true,
		},
		{
			input: 16,
			want:  true,
		},
		{
			input: 218,
			want:  false,
		},
		{
			input: 3,
			want:  false,
		},
		{
			input: 5,
			want:  false,
		},
	}

	for _, tc := range ts {
		got := isPowerOfTwo(tc.input)
		if tc.want != got {
			t.Errorf("input=%d want=%v got=%v", tc.input, tc.want, got)
		}

	}

}
