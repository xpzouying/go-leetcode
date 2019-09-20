package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 方法1: 递归方法
// 递归方法：
// 1. 输出左节点
// 2. 输出自身节点
// 3. 输出右节点
//
// 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗 :2.1 MB, 在所有 Go 提交中击败了96.86%的用户
//
// func inorderTraversal(root *TreeNode) []int {

//     if root == nil {
//         return nil
//     }

//     res := make([]int, 0, 2)

//     doInOrder(root, &res)

//     return res
// }

// func doInOrder(root *TreeNode, list *[]int) {
//     if root == nil {
//         return
//     }

//     if root.Left != nil {
//         doInOrder(root.Left, list)
//     }

//     *list = append(*list, root.Val)

//     if root.Right != nil {
//         doInOrder(root.Right, list)
//     }
// }

// --- 2. 使用非递归的方式遍历 ---
//
// 解题思路：
// 如果使用非递归的方式，需要解决的主要问题是：如何回退到开始计算的上层节点。
// 需要用某种方式来保存。
//
// 这里使用栈来解决
//
// 1. 先押入树/子树的root节点 到 stack
// 2. 如果有左子节点，一直押入左子节点。因为中序遍历中，左子节点的优先级高于root节点
// 3. 如果没有左子节点，则输出root节点值，并且将右子节点押入stack中
//
func inorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)

	stack := make([]*TreeNode, 0)

	curr := root

	for {
		// 如果stack为空，并且没有新的右节点，则结束
		if len(stack) == 0 && curr == nil {
			break
		}

		for {
			if curr == nil {
				break
			}

			stack = append(stack, curr)
			curr = curr.Left
		}

		// 如果没有左子节点可以押入，那么pop栈顶
		curr = stack[len(stack)-1]   // stack.top
		stack = stack[:len(stack)-1] // stack.pop

		// 输出当前节点的值
		res = append(res, curr.Val)

		// 押入右子树
		curr = curr.Right
	}

	return res
}

// --- 3. 颜色标记法 ---
// 推荐：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/yan-se-biao-ji-fa-yi-chong-tong-yong-qie-jian-ming/
// 简单易懂
//
// func inorderTraversal(root *TreeNode) []int {
// }
