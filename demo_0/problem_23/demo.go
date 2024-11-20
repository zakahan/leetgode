package problem_23

/**
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

lists = [[1,4,5],[1,3,4],[2,6]]
[1,1,2,3,4,4,5,6]
维护一个指针数组，

*/

// 本方法为顺序合并 时间复杂度为kN,k代表有k个子list，N表示总长度
func mergeKLists1(lists []*ListNode) *ListNode {
	var startNode = ListNode{0, nil}
	var p = &startNode // 滑动的指针

	// 比较 然后添加

	for true {
		var value = 10001
		var key = 0
		var flag = 0
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && lists[i].Val < value {
				// 首先是这些节点都是非空的，其次是这些节点的值要比value小
				value = lists[i].Val
				key = i
				flag = 1
			}
		}
		if flag == 1 {
			// 比较结束，找到了最佳点
			p.Next = lists[key]
			p = p.Next
			lists[key] = lists[key].Next // 下一步
		} else {
			p.Next = nil
			break // 证明全是nil了
		}

	}
	return startNode.Next
}
