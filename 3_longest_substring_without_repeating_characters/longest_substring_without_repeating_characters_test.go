package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	ts := []struct {
		Input string
		Want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tc := range ts {
		got := lengthOfLongestSubstring(tc.Input)
		assert.Equal(t, tc.Want, got)
	}

}
