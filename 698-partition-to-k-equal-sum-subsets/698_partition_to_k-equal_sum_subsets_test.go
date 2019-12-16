package partition_to_k_equal_sum_subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartitionKSubsets(t *testing.T) {

	ts := []struct {
		Nums []int
		K    int
		Want bool
	}{
		{Nums: []int{2, 2, 2, 2, 3, 4, 5}, K: 4, Want: false},
		{Nums: []int{4, 3, 2, 3, 5, 2, 1}, K: 4, Want: true},
		{Nums: []int{3, 2, 1}, K: 3, Want: false},
	}

	for _, tc := range ts {
		got := canPartitionKSubsets(tc.Nums, tc.K)
		assert.Equal(t, tc.Want, got)
	}

}
