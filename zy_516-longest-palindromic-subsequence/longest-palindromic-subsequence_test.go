package longest_palindromic_subsequence

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {

	ts := []struct {
		input string
		want  int
	}{
		{"a", 1},
		{"", 0},
		{"bbbab", 3},
		{"bb", 2},
		{"cabbad", 4},
		{"cbbcd", 4},
	}

	for _, tc := range ts {
		s := tc.input
		want := tc.want
		got := longestPalindromeSubseq(s)

		if want != got {
			t.Errorf("s=%s, want=%d, but got=%d", s, want, got)
		}
	}

}
