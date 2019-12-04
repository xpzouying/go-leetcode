package diameter_of_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max int
	helper(root, &max)

	return max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	var leftLen int
	var rightLen int

	if root.Left != nil {
		leftLen = helper(root.Left, max) + 1
	}

	if root.Right != nil {
		rightLen = helper(root.Right, max) + 1
	}

	// 更新最长半径
	if leftLen+rightLen > *max {
		*max = leftLen + rightLen
	}

	return MAX(leftLen, rightLen)
}

func MAX(a, b int) int {
	if a > b {
		return a
	}

	return b
}
