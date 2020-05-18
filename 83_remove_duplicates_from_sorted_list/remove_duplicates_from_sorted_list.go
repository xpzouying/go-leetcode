package remove_duplicates_from_sorted_list

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := head
	for {
		if p.Next == nil {
			break
		}

		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
			continue
		}

		p = p.Next
	}

	return head
}

/**

解体思路：

如果下一个节点跟当前节点的值，则删除下一个节点。

删除方法：

将 “当前节点的Next” 指向 “下一个节点的Next”

*/
