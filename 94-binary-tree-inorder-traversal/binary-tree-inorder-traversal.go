package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归方法
func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	res := make([]int, 0, 2)

	doInOrder(root, &res)

	return res
}

func doInOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		doInOrder(root.Left, list)
	}

	*list = append(*list, root.Val)

	if root.Right != nil {
		doInOrder(root.Right, list)
	}
}

/**

递归方法：

1. 输出左节点
2. 输出自身节点
3. 输出右节点

*/
