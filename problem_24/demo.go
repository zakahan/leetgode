package problem_24

/**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

怎么又tm是链表的题目啊？

输入：head = [1,2,3,4]
输出：[2,1,4,3]


[1,2,3,4] => [1,3] [2,4] => [2,1,4,3]  两趟
有什么一趟的方法吗

p->1 q->2
p->next = q->next
q->next = p
then
p = p->next
q = p->next


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil { // 0个节点
		return head
	}
	if head.Next == nil { // 1个节点
		return head
	}
	var startHead ListNode = ListNode{0, head}
	var p *ListNode = head
	var q *ListNode = head.Next
	var z *ListNode = &startHead

	for z.Next != nil && z.Next.Next != nil {
		// swap
		z.Next = q
		p.Next = q.Next
		q.Next = p
		// change
		z = z.Next.Next
		p = z.Next
		if p != nil {
			q = p.Next
		} else {
			break
		}
	}

	return startHead.Next
}
