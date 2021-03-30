package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorGame(t *testing.T) {

	ts := []struct {
		n int

		want bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, true},
		{7, false},
		{8, true},
		{9, false},
		{10, true},
		{11, false},
	}

	for _, tc := range ts {
		got := divisorGame(tc.n)

		assert.Equal(t, tc.want, got)
	}
}
