package balanced_binary_tree

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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if DIFF(treeDeepth(root.Left), treeDeepth(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 返回指定树的高度
func treeDeepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return MAX(treeDeepth(root.Left), treeDeepth(root.Right)) + 1
}

func MAX(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func DIFF(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

/**

解体思路：

判断左子树和右子树的最深深度


*/
