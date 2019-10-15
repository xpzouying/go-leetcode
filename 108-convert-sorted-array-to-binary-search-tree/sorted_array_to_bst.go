package sorted_array_to_bst

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := int(len(nums) / 2)

	root := &TreeNode{Val: nums[mid]}

	// left tree
	if mid != 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}

	// right tree
	if mid < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}

/**

解体思路：

可以使用"递归"的方法。

由于需要做一个"高度平衡的二叉树"，所以可以选择数组中的中位数作为根结点。

1. 取出数组中中位数作为根结点
2. 递归左侧的数组 作为左子树
3. 递归右侧的数组 作为右子树

跳出递归的条件：nums为空

*/
