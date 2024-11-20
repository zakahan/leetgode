package problem_23

/**
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

本次采用分治合并的方式实现

*/

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return MergeTwoLists(lists[0], lists[1])
	}
	// 至少三个的话
	var n = len(lists)
	// 继续用快乐的队列模拟
	queue := []*ListNode{}
	for i := 0; i < n; i++ {
		queue = append(queue, lists[i]) // 先进队列
	}
	var m = len(queue)
	for m > 1 {
		list1 := queue[0]
		list2 := queue[1]
		queue = queue[2:]
		tmpList := MergeTwoLists(list1, list2)
		queue = append(queue, tmpList)
		m = len(queue)
	}

	return queue[0]

}

// 合并两个的
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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
