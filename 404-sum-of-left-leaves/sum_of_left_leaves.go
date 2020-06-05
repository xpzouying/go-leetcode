package sum_of_left_leaves

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

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		// 如果根节点即为空，则返回0
		return 0
	}

	left, right := root.Left, root.Right

	sum := 0
	// 遍历左边子树
	if left != nil {
		if left.Left == nil && left.Right == nil {
			// 如果是左叶子节点
			sum += left.Val
		} else {
			// 否则进行遍历
			sum += sumOfLeftLeaves(left)
		}
	}

	// 遍历右边子树
	if right != nil {
		sum += sumOfLeftLeaves(right)
	}

	return sum
}
