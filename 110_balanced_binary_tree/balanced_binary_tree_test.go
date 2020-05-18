package balanced_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	leftLeef2 := &TreeNode{Val: 4}
	leftLeef := &TreeNode{Val: 3, Left: leftLeef2}
	leftChild := &TreeNode{Val: 2, Left: leftLeef}
	root := &TreeNode{Val: 1, Left: leftChild}

	assert.Equal(t, false, isBalanced(root))
}

func TestIsBalanced2(t *testing.T) {
	leftLeef := &TreeNode{Val: 3}
	leftChild := &TreeNode{Val: 2, Left: leftLeef}
	root := &TreeNode{Val: 1, Left: leftChild}

	assert.Equal(t, false, isBalanced(root))
}
