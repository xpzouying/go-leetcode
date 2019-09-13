package nim_game

import "testing"

func TestNimGame(t *testing.T) {

	ts := []struct {
		input int
		want  bool
	}{
		{
			input: 4,
			want:  false,
		},
		{
			input: 5,
			want:  true,
		},
		{
			input: 6,
			want:  true,
		},
		{
			input: 7,
			want:  true,
		},
	}

	for _, tc := range ts {
		got := canWinNim(tc.input)
		if tc.want != got {
			t.Errorf("input: %v, want: %v, got: %v",
				tc.input, tc.want, got)
		}
	}

}
