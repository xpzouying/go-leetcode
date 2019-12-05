package middle_of_the_linked_list

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

func middleNode(head *ListNode) *ListNode {
	p1 := head // 移动1步
	p2 := head // 移动2步

	for {
		if p2 == nil || p2.Next == nil {
			break
		}

		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1
}
