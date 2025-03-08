// -------------------------------------------------
// Package second_time
// Author: hanzhi
// Date: 2025/3/5
// -------------------------------------------------

package second_time

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	returnList := ListNode{Val: 0, Next: nil}
	z := &returnList
	c := 0

	for p != nil || q != nil {
		var np, nq = 0, 0
		if p != nil {
			np = p.Val
			p = p.Next
		}
		if q != nil {
			nq = q.Val
			q = q.Next
		}
		x := np + nq + c
		//fmt.Println(x)
		c = x / 10
		x = x % 10

		m := ListNode{Val: x, Next: nil}
		z.Next = &m
		z = z.Next
	}
	if c != 0 {
		z.Next = &ListNode{Val: c, Next: nil}
	}

	return returnList.Next
}
