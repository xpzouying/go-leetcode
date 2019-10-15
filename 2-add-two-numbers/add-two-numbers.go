package add_two_numbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	head := new(ListNode)
	tail := head

	// 进位符标志
	carry := 0
	p1 := l1
	p2 := l2

	for {
		if p1 == nil && p2 == nil {
			break
		}

		sum := 0
		if p1 != nil {
			sum += p1.Val
		}
		if p2 != nil {
			sum += p2.Val
		}
		if carry != 0 {
			sum += carry
		}

		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}

		node := &ListNode{
			Val: sum,
		}
		tail.Next = node
		tail = tail.Next

		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	if carry != 0 {
		node := &ListNode{
			Val: 1,
		}
		tail.Next = node
	}

	return head.Next
}

/**

解题思路

由于是逆序排列，所以

*/
