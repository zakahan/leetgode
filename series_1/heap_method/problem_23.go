// -------------------------------------------------
// Package heap_method
// Author: hanzhi
// Date: 2025/2/24
// -------------------------------------------------

package heap_method

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/

func mergeKLists(lists []*ListNode) *ListNode {
	// 遍历，将lists中的每个头节点加入到小顶堆里，每次实际上取出最小值，
	resultNode := ListNode{Val: 0, Next: nil}
	p := &resultNode // 搞个指针
	// 初始化小顶堆
	h := &IHeap{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			// 将内容push进去
			heap.Push(h, lists[i])
		}
	}

	// 不断的入堆，出堆
	for h.Len() > 0 {
		//// 打印
		//fmt.Print("当前的堆：")
		//for _, e := range *h {
		//	fmt.Printf("%v ", e.Val)
		//}
		//fmt.Println()
		// 出堆
		elem := heap.Pop(h)
		elemListNode := elem.(*ListNode)
		p.Next = elemListNode // 取最小值加进去
		p = p.Next

		// 哪个队列出去了，那么就入哪一边的next
		if elemListNode.Next != nil {
			elemListNodeNext := elemListNode.Next
			heap.Push(h, elemListNodeNext) // 入堆
		}

	}

	return resultNode.Next
}

type IHeap []*ListNode

// 求长度的、比较大小的、交换的、堆尾添加元素的、堆头出去的（实际上尾部）

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IHeap) Pop() interface{} {
	old := *h     // 旧有队列
	n := len(old) // 旧队列长度
	x := old[n-1] // 取最后一个元素（实际上，处理过后栈顶在这里）
	*h = old[0 : n-1]
	return x
}
