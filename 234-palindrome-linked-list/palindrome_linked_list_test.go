package palindrome_linked_list

import "testing"

func TestIsPalindrome(t *testing.T) {
	ts := []struct {
		input []int
		want  bool
	}{
		{input: []int{1, 2, 1}, want: true},
		{input: []int{1, 2, 2, 1}, want: true},
		{input: []int{1, 2, 3, 1}, want: false},
	}

	for _, tc := range ts {
		head := newList(tc.input)
		got := isPalindrome(head)
		if tc.want != got {
			t.Errorf("input:%v, want:%v, got:%v", tc.input, tc.want, got)
		}
	}

}

func newList(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, n := range nums {
		p.Next = &ListNode{Val: n}
		p = p.Next
	}

	return head.Next
}
