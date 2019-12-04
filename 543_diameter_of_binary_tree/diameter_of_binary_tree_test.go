package diameter_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tree := &TreeNode{Val: 1}
	got := diameterOfBinaryTree(tree)
	assert.Equal(t, 0, got)
}
