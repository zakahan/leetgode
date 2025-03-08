package problem_206

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil { // 至少得有一个
		return head // //
	}
	// 头插法
	fNode := ListNode{0, nil}
	q := head
	p := head.Next
	for p != nil {
		q.Next = fNode.Next // 修改q的next
		fNode.Next = q

		q = p
		p = p.Next
	}
	if q != nil {
		q.Next = fNode.Next
		fNode.Next = q
	}

	return fNode.Next
}
