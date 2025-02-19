// -------------------------------------------------
// Package problem_300
// Author: hanzhi
// Date: 2025/2/11
// -------------------------------------------------

package problem_300

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	d := make([]int, 0)
	d = append(d, nums[0])
	// 我们维护一个数组 d[i] ，表示长度为 i 的最长上升子序列的末尾元素的最小值
	for i := 1; i < len(nums); i++ {
		fmt.Println(d)
		if nums[i] > d[len(d)-1] {
			// 如果增长，那么很好，加进去即可
			d = append(d, nums[i])
		} else if nums[i] == d[len(d)-1] {
			// 如果等于，直接跳过，没有任何意义
		} else {
			// 如果小于，那么要二分查找替代掉那个比他小的值
			// 用二分查找定位，然后再替换掉
			loc := bSearch(d, 0, len(d)-1, nums[i])
			d[loc] = nums[i]
		}
	}

	return len(d)
}

func bSearch(nums []int, start, end int, target int) int {
	if start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= target {
			// target在左侧
			return bSearch(nums, start, mid-1, target)
		} else {
			return bSearch(nums, mid+1, end, target)
		}
	} else {
		return start
	}
}
