package problem_19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 倒数第n个，双指针
	// head -> 5 4 3 2 1 nil
	// 删除倒数第二个 先走两步，q=node_4
	// 然后一起走,直到q.Next=nil head-4 ; 5-3 ; 4-2 ; 3-1 ; 删除3.next
	//
	var headNode ListNode = ListNode{0, head}
	var p = &headNode
	var q = &headNode
	var i int = 0
	// 删除就是删除 p的next
	for i < n && p != nil {
		p = p.Next // 不考虑越界问题
		i += 1
	}
	for p.Next != nil {
		q = q.Next
		p = p.Next
	}
	p = q.Next
	q.Next = p.Next
	return headNode.Next
}
