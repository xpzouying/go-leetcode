package diameter_of_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var max int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	helper(root)

	return max
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	var leftLen int
	var rightLen int

	if root.Left != nil {
		leftLen = helper(root.Left) + 1
	}

	if root.Right != nil {
		rightLen = helper(root.Right) + 1
	}

	// 更新最长半径
	max = MAX(max, leftLen+rightLen)

	return MAX(leftLen, rightLen)
}

func MAX(a, b int) int {
	if a > b {
		return a
	}

	return b
}
