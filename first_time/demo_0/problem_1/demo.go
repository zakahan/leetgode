package _1_two_sum

import (
	"sort"
)

func TwoSum_n2(nums []int, target int) []int {
	// 先写一个最简单的，就是暴力搜索的方式
	for i := 0; i < len(nums)-1; i++ {
		x := nums[i]
		for j := i + 1; j < len(nums); j++ {
			y := nums[j]
			if x+y == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSum_nlogn(nums []int, target int) []int {
	// 先排序 后找坐标
	argNum := ArgSort(nums)
	sort.Ints(nums)
	var i = 0
	var j = len(nums) - 1
	for i != j {
		if nums[i]+nums[j] == target {
			return []int{argNum[i], argNum[j]}
		} else if nums[i]+nums[j] > target {
			// 说明需要缩小
			j--

		} else {
			i++
		}

	}
	return nil
}

func ArgSort(arr []int) []int {
	// 创建索引切片
	indices := make([]int, len(arr))
	for i := range indices {
		indices[i] = i
	}

	// 根据sort进行排序
	sort.Slice(indices, func(i int, j int) bool {
		return arr[indices[i]] < arr[indices[j]]
	})

	return indices
}

func TwoSum_n(nums []int, target int) []int {
	// 建立一个哈希表， 值是索引， 键是nums的数值
	myMap := make(map[int]int)
	for i, item := range nums {
		if value, ok := myMap[target-item]; ok {
			return []int{i, value}
		}
		myMap[item] = i
	}
	return nil
}
