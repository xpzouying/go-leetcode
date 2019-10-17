package palindrome_linked_list

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

func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	// 1. 找到中间节点
	p := head // 慢指针
	q := head // 快指针

	// 是否是奇数个节点
	odd := false

	for {

		if q == nil {
			// 如果是偶数个节点，则q为空
			odd = false
			break
		}

		if q != nil && q.Next == nil {
			// 如果是奇数个节点，则q不为空，q的next为空
			odd = true
			break
		}

		p = p.Next
		q = q.Next.Next
	}

	// 后半段链表，second为后半段链表的第一个
	var second *ListNode
	// 如果是奇数
	if odd {
		second = p.Next
	} else {
		second = p
	}

	// 2. 反转前半段链表
	first := &ListNode{}
	curr := head
	for {
		if curr == p {
			break
		}

		n := curr.Next
		curr.Next = first.Next
		first.Next = curr
		curr = n
	}
	first = first.Next

	// 如果是偶数
	return doCheck(first, second)
}

// 检查两个列表是否相同
func doCheck(l1, l2 *ListNode) bool {
	for {
		if l1 == nil && l2 == nil {
			return true
		}

		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}
}
