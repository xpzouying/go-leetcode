package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}

	head := new(ListNode)

	p := head

	for _, i := range l {
		n := &ListNode{
			Val: i,
		}

		p.Next = n

		p = p.Next
	}

	return head.Next
}

func (l *ListNode) Lists() []int {
	var ss []int

	for {
		if l == nil {
			break
		}

		ss = append(ss, l.Val)

		l = l.Next
	}

	return ss
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	var p *ListNode

	l3 = &ListNode{Next: nil}
	p = l3

	i := l1
	j := l2

	for {
		if i == nil || j == nil {
			break
		}

		if i.Val < j.Val {
			p.Next = i
			i = i.Next
		} else {
			p.Next = j
			j = j.Next
		}

		p = p.Next
	}

	if i != nil {
		p.Next = i
	}
	if j != nil {
		p.Next = j
	}

	return l3.Next
}

func equalList(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}
