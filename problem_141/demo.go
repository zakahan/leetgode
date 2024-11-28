package problem_141

/*
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// 我记得好像是双指针就可以，前面走两步，后面走一步，如果发现追上了那就说明是有重复(官方叫龟兔赛跑解，我觉得可以叫做悬魂梯解法，让我想起来龙岭迷窟悬魂梯那章）
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return false
		} else {
			p2 = p2.Next
		}
		if p2 == nil {
			return false
		}
		if p2 == p1 || p2.Next == p1 {
			return true
		}
	}
	return false
}
