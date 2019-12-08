package threesum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	ts := []struct {
		Nums []int
		Want [][]int
	}{
		{
			Nums: []int{-1, 0, 1, 2, -1, -4},
			Want: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		}, {
			Nums: []int{0, 0, 0, 0},
			Want: [][]int{
				[]int{0, 0, 0},
			},
		},
	}

	for _, tc := range ts {
		got := threeSum(tc.Nums)
		assert.Equal(t, tc.Want, got)
	}
}
