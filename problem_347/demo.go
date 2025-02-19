// -------------------------------------------------
// Package problem_347
// Author: hanzhi
// Date: 2025/2/19
// -------------------------------------------------

package problem_347

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h) // 交换和下浮等过程由container实现了

		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }            // 返回堆的大小。
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] } // 定义堆中元素的比较规则。
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }  // 交换堆中两个元素的位置。

func (h *IHeap) Push(x interface{}) { // 向堆中添加新元素。（堆尾）
	*h = append(*h, x.([2]int))
}

// Pop Pop实际上是移除堆顶，但我需要实现的不是这样，我需要实现的是移除最后一个，因为接口在这之前已经通过Swap交换了。
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
