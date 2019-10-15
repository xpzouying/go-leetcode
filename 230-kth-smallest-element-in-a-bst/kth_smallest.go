package kth_smallest

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

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	leftNodesCount := countNodes(root.Left)

	// 如果是根节点
	if leftNodesCount+1 == k {
		return root.Val
	}

	// 如果左子树的节点不够k，那么搜索右边的子树
	if leftNodesCount+1 < k {
		return kthSmallest(root.Right, k-leftNodesCount-1)
	}

	return kthSmallest(root.Left, k)
}

// RETURN NODE COUNT
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 1 for root
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

/**

解题思路：

由于二叉搜索树是：左子树小于跟节点， 跟节点小于右子树

1. 统计左节点的个数，有N个节点

2. 若K>N+1：则在右子树，搜索 K-（N+1）

3. 若K<N+1：则在左子树，搜索K

4。K==N+1: 则返回跟节点

*/
