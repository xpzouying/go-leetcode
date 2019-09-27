package convert_bst_to_greater_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	helper(root, 0)

	return root
}

func helper(root *TreeNode, sum int) int {

	if root == nil {
		return sum
	}

	// 计算出所有大于当前节点的累计和
	root.Val += helper(root.Right, sum)

	// 迭代左子树，左子树的初始累加值为父节点的值
	return helper(root.Left, root.Val)
}

/**

解题思路：

题目中可以得出几点：

1. 二叉搜索树：可以得出左子树比根节点小，右子树比根节点大的规律。这个是二叉搜索树重要的解题线索。

2. 要把所有大于它节点值的和，转换成累加树。可以得出中序遍历得到递增的结果。
如果是反中序，即先右子树，再根节点，再左子树，即可得到累加树。

那么可以推出解法：

1. 使用递归算法。

2. 先遍历右子树，累计出所有大于当前节点的值。

3. 再遍历左子树。


*/
