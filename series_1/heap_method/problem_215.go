// -------------------------------------------------
// Package heap_method
// Author: hanzhi
// Date: 2025/2/24
// -------------------------------------------------

package heap_method

import (
	"container/heap"
	"fmt"
	"math"
)

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

// 这道题用这种方法肯定是不够o(n)的，但是 用来练手可以
// 先选取K个构建堆，然后扔掉最小值，再来，再仍，一直到全部nums遍历完成，就剩下K个最大值了，这里面，堆顶就是第k大
// 建堆是o(k) 然后调整o(logk) * (n-k)次
func findKthLargest(nums []int, k int) int {
	h := PHeap{}
	heap.Init(&h)
	maxValue := math.MinInt32
	for i := 0; i < k; i++ {
		heap.Push(&h, nums[i])
	}
	// 出入
	for i := k; i < len(nums); i++ {
		fmt.Println(h)
		value := heap.Pop(&h)
		if value.(int) > maxValue {
			maxValue = value.(int)
		}
		heap.Push(&h, nums[i])
	}
	value := heap.Pop(&h)
	if maxValue > value.(int) {
		return maxValue
	}
	return value.(int)

}

type PHeap []int

// swap， less， len pop push

func (h PHeap) Len() int {
	return len(h)
}

func (h PHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h PHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *PHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
