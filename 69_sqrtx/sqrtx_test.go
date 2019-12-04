package sqrtx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	ts := []struct {
		number int
		want   int
	}{
		{number: 4, want: 2},
		{number: 8, want: 2},
		{number: 9, want: 3},
		{number: 11, want: 3},
	}

	for _, tc := range ts {
		got := mySqrt(tc.number)
		assert.Equal(t, tc.want, got)
	}

}
