package invert_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    doInvertTree(root)
    
    return root
}


func doInvertTree(root *TreeNode) {
    if root == nil {
        return
    }
    
    left := root.Left
    right := root.Right
    
    root.Left = right
    root.Right = left
    
    
    doInvertTree(root.Left)
    doInvertTree(root.Right)
}