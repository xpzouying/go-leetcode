package binary_tree_paths

import "testing"

func TestBinaryTreePaths(t *testing.T) {

	ts := []struct {
		Root *TreeNode
		Want []string
	}{
		{
			Root: nil,
			Want: []string{},
		},
		{
			Root: &TreeNode{Val: 0},
			Want: []string{"0"},
		},
		{
			Root: newTree1(),
			Want: []string{"1->2->5", "1->3"},
		},
	}

	for _, tc := range ts {
		got := binaryTreePaths(tc.Root)

		want := tc.Want
		if len(want) != len(got) {
			t.Errorf("got:%+v not equal want:%+v", got, want)
		}

		l := len(want)
		for i := 0; i < l; i++ {
			if want[i] != got[i] {
				t.Fatalf("got:%+v not equal want:%+v", got, want)
			}
		}
	}
}

func newTree1() *TreeNode {

	root := &TreeNode{Val: 1}

	rootLeft := &TreeNode{Val: 2}
	rootRight := &TreeNode{Val: 3}
	root.Left = rootLeft
	root.Right = rootRight

	rootLeft.Right = &TreeNode{Val: 5}

	return root
}
