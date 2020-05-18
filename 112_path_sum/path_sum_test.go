package path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}

	root := &TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}

	assert.Equal(t, false, hasPathSum(root, 1))
	assert.Equal(t, false, hasPathSum(root, 2))
	assert.Equal(t, true, hasPathSum(root, 3))
	assert.Equal(t, true, hasPathSum(root, 4))
}

func TestHasPathSum2(t *testing.T) {
	right := &TreeNode{Val: -3}

	root := &TreeNode{
		Val:   -2,
		Right: right,
	}

	assert.Equal(t, true, hasPathSum(root, -5))
}
