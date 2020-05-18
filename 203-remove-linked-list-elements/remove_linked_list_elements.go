package remove_linked_list_elements

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	p := &ListNode{
		Next: head,
	}

	move := p
	for {
		if move.Next == nil {
			break
		}

		// 如果下一个节点值为val
		if move.Next.Val == val {
			move.Next = move.Next.Next
			continue
		}

		move = move.Next
	}

	return p.Next
}
