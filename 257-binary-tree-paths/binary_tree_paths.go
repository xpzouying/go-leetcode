package binary_tree_paths

import "fmt"

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

func binaryTreePaths(root *TreeNode) []string {

	res := make([]string, 0)

	if root == nil {
		return res
	}

	helper(root, "", &res)

	return res
}

// 遍历到当前节点时，进行递归。
// path: 遍历到当前节点时，path的结果
func helper(node *TreeNode, path string, res *[]string) {

	// 添加当前节点的值
	if path == "" {
		// 如果是第一个节点
		path = fmt.Sprintf("%d", node.Val)
	} else {
		path = fmt.Sprintf("%s->%d", path, node.Val)
	}

	// 判断是否还有子节点，如果有的话，就表示当前节点不是叶子节点
	// 否则，当前节点是叶子节点，记录结果后退出
	if node.Left == nil && node.Right == nil {
		// 如果已经到达了叶子节点
		*res = append(*res, path)

		return
	}

	if node.Left != nil {
		helper(node.Left, path, res)
	}

	if node.Right != nil {
		helper(node.Right, path, res)
	}
}
