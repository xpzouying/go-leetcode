package decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {

	ts := []struct {
		input string
		want  int
	}{
		{
			input: "12",
			want:  2,
		},
		{
			input: "226",
			want:  3,
		},
		{
			input: "0",
			want:  0,
		},
		{
			input: "1",
			want:  1,
		},
		{
			input: "02",
			want:  0,
		},
		{
			input: "27",
			want:  1,
		},
		{
			input: "01",
			want:  0,
		},
		{
			input: "100",
			want:  0,
		},
		{
			input: "101",
			want:  1,
		},
	}

	for _, tc := range ts {

		got := numDecodings(tc.input)
		if tc.want != got {
			t.Errorf("input:%v, want:%d, got:%d",
				tc.input, tc.want, got)
		}
	}

}
