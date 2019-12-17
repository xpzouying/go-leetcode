package merge_k_sorted_lists

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

func mergeKLists(lists []*ListNode) *ListNode {
	ll := len(lists)

	if ll == 0 {
		return nil
	}

	if ll == 1 {
		return lists[0]
	}

	mid := ll / 2

	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeTwo(left, right)
}

func mergeTwo(l1, l2 *ListNode) *ListNode {

	n1 := l1
	n2 := l2

	head := new(ListNode)
	p := head

	for {
		if n1 == nil || n2 == nil {
			break
		}

		if n1.Val < n2.Val {
			p.Next = n1
			n1 = n1.Next
		} else {
			p.Next = n2
			n2 = n2.Next
		}

		p = p.Next
	}

	if n1 != nil {
		p.Next = n1
	}

	if n2 != nil {
		p.Next = n2
	}

	return head.Next
}

/**
解题思路：

大体流程：做分治合并；

时间复杂度：

由于列表数组长度为K，使用分治合并后的算法为：log(K)，
每次合并的长度是n，
所以总体时间复杂度为：n * log(K)

空间复杂度为：最终列表长度：o(n)
*/
