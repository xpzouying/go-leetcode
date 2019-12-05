package merge_two_binary_trees

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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	newNode := mergeCurrentNode(t1, t2)
	newNode.Left = mergeTrees(t1.Left, t2.Left)
	newNode.Right = mergeTrees(t1.Right, t2.Right)

	return newNode
}

func mergeCurrentNode(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	n := &TreeNode{}
	if t1 == nil && t2 != nil {
		n.Val = t2.Val
	} else if t1 != nil && t2 == nil {
		n.Val = t1.Val
	} else {
		n.Val = t1.Val + t2.Val
	}

	return n
}
