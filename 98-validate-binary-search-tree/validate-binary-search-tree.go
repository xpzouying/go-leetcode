package validate_binary_search_tree

import "math"

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

/**

解题思路：

判断是否是查找二叉树时，不光是需要检查左右子节点是否满足，
还要判断左右子树是否满足。

所以在递归时，需要输入取值的范围。
*/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, math.MinInt64, root.Val) && check(root.Right, root.Val, math.MaxInt64)
}

// check检查root子树是否满足查找二叉树的条件
// minVal：最小取值
// maxVal：最大取值
func check(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {

		leftMax := MIN(maxVal, root.Val)

		if check(root.Left, minVal, leftMax) == false {
			return false
		}
	}

	if root.Right != nil {

		rightMin := MAX(root.Val, minVal)

		if check(root.Right, rightMin, maxVal) == false {
			return false
		}
	}

	return true
}

func MAX(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MIN(a, b int) int {
	if a < b {
		return a
	}

	return b
}
