package problem_31

import "sort"

// 从最后一位开始找i和j
// 如果 i 和 i后面的东西比(j)  如果n[i] < n[j] 交换并对i+1:end的内容做排序

// 如果都不是，反转

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				// 让j到i之间最小（包括i不包括j）
				sort.Slice(nums[i+1:], func(a, b int) bool {
					return nums[i+1+a] < nums[i+1+b]
				})
				return
			}

		}
	}

	var i, j int
	i = 0
	j = len(nums) - 1
	// 反转
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return

}
