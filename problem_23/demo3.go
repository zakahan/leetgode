package problem_23

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

本次采用最小堆（优先队列）实现
算法的逻辑
假设一共有k个链表，那么第一次的“潜在最小值”一定在k个链表的头节点中寻找，
选出了一个投，潜在最小值就是其他k-1个头节点 再加上选出去的那条的二号节点
这样就转化为了一个多次重复的topK问题，topK问题正是堆擅长的
*/
type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
} // 最小堆
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v any) {
	*h = append(*h, v.(*ListNode))
}
func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

//func mergeKLists(lists []*ListNode) *ListNode {
//    h := hp{}
//}
