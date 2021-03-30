package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCuttingRope(t *testing.T) {
	ts := []struct {
		n int

		want int
	}{
		{2, 1},
		{10, 36},
	}

	for _, tc := range ts {
		got := cuttingRope(tc.n)

		assert.Equal(t, tc.want, got,
			"n=%d,want=%d", tc.n, tc.want)
	}
}
