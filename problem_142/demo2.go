package problem_142

/*
来分析一下快慢指针

超过边界的好说，那种不用算

fast 指针每次走两步
slow指针每次走一步

假设 外面长度a，圈内长度b， 入口坐标是x，相遇点是y
那么slow走了x+y
fast走了 2(x+y) = x + y + (x+y)

故而，n圈等于 y + x + y
让slow再走 x+y 就能回到相遇点，那么也就是再走x就到了入口！
但是我不知道x是多少，只能凑了，新整一个head开始的节点出发，都走x，会在入口相遇



*/

func detectCycle(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	if head == nil {
		return nil
	}
	for {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}
		if p2 == nil || p2.Next == nil {
			return nil
		}
		if p2 == p1 {
			q := head
			for q != p1 {
				p1 = p1.Next
				q = q.Next
			}
			return q
		}

	}
}
