package sum_of_left_leaves

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {

	ts := []struct {
		Tree *TreeNode
		Want int
	}{
		{Tree: newTree1(), Want: 24},
		{Tree: newTree2(), Want: 9},
	}

	for _, tc := range ts {
		tree := tc.Tree
		want := tc.Want

		got := sumOfLeftLeaves(tree)
		if want != got {
			t.Errorf("want=%d, buf got=%d", want, got)
		}
	}

}

func newTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 9}
	right := &TreeNode{Val: 20}

	root.Left = left
	root.Right = right

	right.Left = &TreeNode{Val: 15}
	right.Right = &TreeNode{Val: 7}

	return root
}

func newTree2() *TreeNode {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 9}
	right := &TreeNode{Val: 20}

	root.Left = left
	root.Right = right

	right.Right = &TreeNode{Val: 7}

	return root
}
