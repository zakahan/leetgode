// -------------------------------------------------
// Package problem_287
// Author: hanzhi
// Date: 2025/1/31
// -------------------------------------------------

package problem_287

func findDuplicate(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		var mid = (start + end) / 2
		var count = 0
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		// 遍历完成，如果count不等于mid个，说明在左边，否则在右边
		if count > mid {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}
