package two_sum_ii_input_array_is_sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	ts := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 2},
		},
		{
			nums:   []int{2, 7, 11, 15},
			target: 18,
			want:   []int{2, 3},
		},
	}

	for _, tc := range ts {
		got := twoSum(tc.nums, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
