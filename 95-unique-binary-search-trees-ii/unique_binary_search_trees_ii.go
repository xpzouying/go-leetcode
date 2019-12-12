package unique_binary_search_trees_ii

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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return doGenerateTrees(1, n)
}

// 递归生成从[left, right]区间范围内的树
func doGenerateTrees(left, right int) []*TreeNode {

	if left > right {
		return []*TreeNode{nil}
	}

	if left == right {
		return []*TreeNode{&TreeNode{Val: left}}
	}

	trees := make([]*TreeNode, 0, right-left+1)
	for i := left; i <= right; i++ {
		// 以i为根节点，生成左右子节点

		leftTrees := doGenerateTrees(left, i-1)
		rightTrees := doGenerateTrees(i+1, right)

		// 拼接所有的树
		for m := 0; m < len(leftTrees); m++ {
			for n := 0; n < len(rightTrees); n++ {
				tree := &TreeNode{Val: i}

				tree.Left = leftTrees[m]
				tree.Right = rightTrees[n]
				trees = append(trees, tree)
			}
		}
	}

	return trees
}
