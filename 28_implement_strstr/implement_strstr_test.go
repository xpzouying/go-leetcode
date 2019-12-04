package implement_strstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	ts := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"1", "2", -1},
		{"1", "22", -1},
	}

	for _, tc := range ts {
		got := strStr(tc.haystack, tc.needle)
		assert.Equal(t, tc.want, got)
	}
}
