package lowest_common_ancestor_of_a_binary_tree

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 其中一个节点在根节点，另外一个节点不需要查找，
	// 因为从根节点开始，必然在其子节点
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		// 左右节点分别找到其中一个节点
		return root
	}

	// 否则返回找到的节点
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}

	return nil
}
