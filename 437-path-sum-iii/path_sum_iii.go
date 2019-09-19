package main

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

func pathSum(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	return doFind(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// 从当前节点开始，以当前节点为路径开始，查找出sum的路径数量
func doFind(node *TreeNode, sum int) int {

	if node == nil {
		return 0
	}

	count := 0

	diff := sum - node.Val
	if diff == 0 {
		count++
	}

	count += doFind(node.Left, diff)
	count += doFind(node.Right, diff)

	return count
}

// 题解 / 思路
//
// 由于可以从任意节点开始，
// 那么我们需要在当前节点做：

// 由于结果不要求以根节点为开始节点，那么我们可以想像成：
// 1. 从当前节点开始，统计满足sum的总数；
// 2. 从左子节点开始，统计满足sum的总数；
// 3. 从右子节点开始，统计满足sum的总数；
// 4. 累加上述3种情况

// 相当于两种类型的累计：
//
// 1. 从当前节点开始，进行计算路径值
// 2. 从当前节点开始统计满足sum的路径

// 统计时的技巧：
// 1. 如果当前节点的值val为需要累计的sum，那么累计+1
// 2. 否则，以当前节点为开始节点，向下查找sum-node.val：
// 2.1 左子节点
// 2.2 右子节点

// 使用递归方法，递归的跳出条件是：root为空
