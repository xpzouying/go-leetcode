package longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {

	ts := []struct {
		Input string
		Want  []string // 可能会有多个答案，只要取出最长的即可
	}{
		{
			Input: "babad",
			Want:  []string{"bab", "aba"},
		},
		{
			Input: "cbbd",
			Want:  []string{"bb"},
		},
		{
			Input: "cccc",
			Want:  []string{"cccc"},
		},
		{
			Input: "ccc",
			Want:  []string{"ccc"},
		},
	}

	for _, tc := range ts {
		got := longestPalindrome(tc.Input)

		if !Contains(tc.Want, got) {
			t.Errorf("input: %s, want: %s, got: %s", tc.Input, tc.Want, got)
		}
	}

}

func Contains(l []string, s string) bool {
	for _, e := range l {
		if s == e {
			return true
		}
	}

	return false
}
