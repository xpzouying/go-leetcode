package symmetric_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return checkLeftAndRight(root.Left, root.Right)
}

func checkLeftAndRight(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	// 只有一侧为空
	if left == nil || right == nil {
		return false
	}

	// 如果左右根节点不同
	if left.Val != right.Val {
		return false
	}

	return checkLeftAndRight(left.Left, right.Right) && checkLeftAndRight(left.Right, right.Left)
}

/**




 */
