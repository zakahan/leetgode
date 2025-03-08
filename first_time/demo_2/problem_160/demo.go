package problem_160

/*
相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表解法 时空复杂度：O(m+n),O(m+n)
func getIntersectionNodeH(headA, headB *ListNode) *ListNode {
	myMap := make(map[*ListNode]int)
	q := headA
	p := headB
	for q != nil {
		myMap[q] = q.Val
		q = q.Next
	}
	for p != nil {
		if myMap[p] != 0 {
			return p
		}
		p = p.Next
	}
	return nil
}

// 双指针法 两趟
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	q := headA
	p := headB
	for q != nil && p != nil {
		q = q.Next
		p = p.Next

	}
	if q == nil {
		// 那么就是p更长
		p2 := headB
		q2 := headA
		for p != nil {
			p = p.Next
			p2 = p2.Next
		}
		for p2 != nil {
			if p2 == q2 {
				return p2
			}
			p2 = p2.Next
			q2 = q2.Next
		}
	}
	if p == nil {
		// 那么就是p更长
		p2 := headB
		q2 := headA
		for q != nil {
			q = q.Next
			q2 = q2.Next
		}
		for q2 != nil {
			if p2 == q2 {
				return p2
			}
			p2 = p2.Next
			q2 = q2.Next
		}
	}
	return nil
}
