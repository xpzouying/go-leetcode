package path_sum

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

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	// 如果是叶子节点
	if root.Right == nil && root.Left == nil {
		// 当前节点正好跟sum相等
		if root.Val == sum {
			return true
		}

		return false
	}

	// 已经超出sum值，已经爆满了
	// NOTE: 题目中带有负数，则没有爆满的情况。
	// if root.Val > sum {
	// 	return false
	// }

	if root.Left != nil && hasPathSum(root.Left, sum-root.Val) {
		return true
	}

	if root.Right != nil && hasPathSum(root.Right, sum-root.Val) {
		return true
	}

	return false
}

/*

解题思路：

由于是判断从根节点到叶子节点的路径，我们可以得知：

1. 从根节点开始，是连续的；
2. 当遍历到某个节点时，累计和会出现下列情况：
  2.1 正好等于SUM，且左右子节点都为空（为叶子节点），则返回true。
  2.2 大于SUM，则返回当前路径为false。
  2.3 小于SUM，则继续向下遍历，遍历到叶子节点为止，继续第2步。

*/
