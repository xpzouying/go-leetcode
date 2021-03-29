package search_a_2d_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {

	ts := []struct {
		matrix [][]int
		target int

		want bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			13,
			false,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50}},
			5,
			true,
		},
	}

	for _, tc := range ts {

		got := searchMatrix(tc.matrix, tc.target)

		assert.Equal(t, tc.want, got,
			"matrix: %+v, target: %d", tc.matrix, tc.target)
	}

}
