package problem_148

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

func sortList(head *ListNode) *ListNode {
	return linkSort(head, nil)
}

func linkSort(head *ListNode, tail *ListNode) *ListNode { // tail就是尾巴
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow := head
	fast := head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(linkSort(head, mid), linkSort(mid, tail))
}

func merge(link1 *ListNode, link2 *ListNode) *ListNode {
	var p *ListNode = link1
	var q *ListNode = link2

	var m ListNode = ListNode{}
	var z *ListNode = &m

	for p != nil && q != nil {
		if p.Val <= q.Val {
			z.Next = p
			p = p.Next
		} else {
			z.Next = q
			q = q.Next
		}
		z = z.Next
	}
	if p != nil {
		z.Next = p
	}
	if q != nil {
		z.Next = q
	}
	return m.Next
}
