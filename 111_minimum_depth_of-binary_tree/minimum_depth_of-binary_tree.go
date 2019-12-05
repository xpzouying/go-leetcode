package minimum_depth_of_binary_tree

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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}

	var left int
	if root.Left != nil {
		left = minDepth(root.Left) + 1
	}

	var right int
	if root.Right != nil {
		right = minDepth(root.Right) + 1
	}

	// 只有左子树
	if left != 0 && right == 0 {
		return left
	}

	// 只有右子树
	if right != 0 && left == 0 {
		return right
	}

	return MIN(left, right)
}

func MIN(a, b int) int {
	if a < b {
		return a
	}

	return b
}
