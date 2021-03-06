package reverse_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {

	ts := []struct {
		Input string
		Want  string
	}{
		{
			Input: "the sky is blue",
			Want:  "blue is sky the",
		},
		{
			Input: "  hello world!  ",
			Want:  "world! hello",
		},
		{
			Input: "a good   example",
			Want:  "example good a",
		},
	}

	for _, tc := range ts {
		got := reverseWords(tc.Input)
		assert.Equal(t, tc.Want, got)
	}
}
