package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {

	ts := []struct {
		digits string

		want []string
	}{

		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for _, tc := range ts {
		digits := tc.digits

		got := letterCombinations(digits)

		sort.Strings(got)

		want := tc.want
		sort.Strings(want)
		assert.Equal(t, want, got)
	}
}
