package problem_21

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 为了空间复杂度尽可能低，我还是得原地操作最好

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var totalNode ListNode = ListNode{0, nil}
	var p *ListNode = list1
	var q *ListNode = list2
	var t *ListNode = &totalNode
	for p != nil || q != nil {
		if q == nil {
			t.Next = p
			break
		}
		if p == nil {
			t.Next = q
			break
		}
		if p.Val < q.Val {
			t.Next = p
			p = p.Next
			t = t.Next
		} else {
			t.Next = q
			q = q.Next
			t = t.Next
		}

	}
	return totalNode.Next
}
