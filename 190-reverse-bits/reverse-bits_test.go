package reverse_bits

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	ts := []struct {
		input uint32
		want  uint32
	}{
		{input: 43261596, want: 964176192},
		{input: 0, want: 0},
		{input: 1, want: 2147483648},
		{input: 2147483648, want: 1},
		{input: 4294967293, want: 3221225471},
	}

	for _, tc := range ts {
		got := reverseBits(tc.input)
		if tc.want != got {
			t.Errorf("reverse bits failed: input:%d (%b), want:%d (%b), got:%d (%b)",
				tc.input, tc.input,
				tc.want, tc.want,
				got, got,
			)
		}
	}

}
