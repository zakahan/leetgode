// -------------------------------------------------
// Package problem_322
// Author: hanzhi
// Date: 2025/2/12
// -------------------------------------------------

package problem_322

import (
	"math"
)

// 这个策略宣告失败了，太慢了。 [1,2,5] 100这种就过不去了
var myMap map[int]int
var globalDeepth int

func coinChange0(coins []int, amount int) int {
	quickSort(coins, 0, len(coins)-1)
	globalDeepth = math.MaxInt64 // 搞个贼大的值
	myMap = make(map[int]int)
	// 层序遍历+剪枝 但是怎么终止呢？
	backTrack(coins, amount, 0)
	if globalDeepth == math.MaxInt64 {
		return -1
	} else {
		return globalDeepth
	}
}

func backTrack(coins []int, target int, deepth int) {
	if target < 0 {
		// 代表这一层失败了
		return
	} else if target == 0 {
		if deepth < globalDeepth {
			globalDeepth = deepth // 修改最小深度
		}
		return
	} else { // 如果target >0说明还能继续
		if myMap[target] != 0 { // 剪枝
			return
		} else {
			myMap[target] = 1
		}
		// 逐个算
		for i := len(coins) - 1; i >= 0; i-- {
			backTrack(coins, target-coins[i], deepth+1)
		}
	}
}

// quickSort 是快速排序的主函数，接受切片、起始索引和结束索引作为参数
func quickSort(arr []int, low, high int) {
	if low < high {
		// 找到分区点
		pivotIndex := partition(arr, low, high)
		// 对左半部分递归排序
		quickSort(arr, low, pivotIndex-1)
		// 对右半部分递归排序
		quickSort(arr, pivotIndex+1, high)
	}
}

// partition 函数用于划分数组，并返回基准值的最终位置
func partition(arr []int, low, high int) int {
	// 选择最后一个元素作为基准值
	pivot := arr[high]
	i := low - 1 // i 是小于 pivot 的区域的最后一个索引

	// 遍历数组，将小于 pivot 的元素放到左侧
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			// 交换 arr[i] 和 arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将 pivot 放到正确的位置
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // 返回 pivot 的索引
}
