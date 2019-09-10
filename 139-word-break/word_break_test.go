package word_break

import "testing"

func TestWordBreak(t *testing.T) {

	ts := []struct {
		S    string
		Dict []string
		Want bool
	}{
		{
			S:    "leetcode",
			Dict: []string{"leet", "code"},
			Want: true,
		},
		{
			S:    "applepenapple",
			Dict: []string{"apple", "pen"},
			Want: true,
		},
		{
			S:    "catsandog",
			Dict: []string{"cats", "dog", "sand", "and", "cat"},
			Want: false,
		},
		{
			S:    "aaaaaaa",
			Dict: []string{"aaa", "aaaa"},
			Want: true,
		},
	}

	for _, tc := range ts {
		got := wordBreak(tc.S, tc.Dict)
		if tc.Want != got {
			t.Errorf("s: %s, dict: %+v, want: %v, got: %v",
				tc.S, tc.Dict, tc.Want, got)
		}
	}

}
