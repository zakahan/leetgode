// -------------------------------------------------
// Package problem_41
// Author: hanzhi
// Date: 2025/2/26
// -------------------------------------------------

package problem_41

import "math"

// 原地处理为哈希

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)

	temp := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 { // 非正整数的情况，直接忽视

		} else { // 正数，能处理
			temp = nums[i]                                            // 将当前值赋值到temp,开始处理
			for temp < len(nums) && temp >= 0 && nums[temp] != temp { // 就说判断，只要temp的值属于len(nums)范围内，那么就粘贴过去
				// 但是要考虑一个终止情况，就说nums[index]=index 说明放对了，那就停止就行了
				// 将temp赋值到nums[index]中
				nums[temp], temp = temp, nums[temp]
			}
			// 直到temp的值超过了范围（超过范围就无所谓了，直接丢弃就行

		}

	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == i {
			// pass
		} else {
			return i
		}
	}
	return len(nums)
}
