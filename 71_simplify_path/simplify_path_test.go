package simplifypath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {

	ts := []struct {
		Input string
		Want  string
	}{
		{
			Input: "/home/",
			Want:  "/home",
		},
		{
			Input: "/a/../../b/../c//.//",
			Want:  "/c",
		},
		{
			Input: "/a/./b/../../c/",
			Want:  "/c",
		},
		{
			Input: "/a//b////c/d//././/..",
			Want:  "/a/b/c",
		},
	}

	for _, tc := range ts {
		got := simplifyPath(tc.Input)
		assert.Equal(t, tc.Want, got)
	}

}
