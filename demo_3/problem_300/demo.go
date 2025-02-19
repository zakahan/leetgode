// -------------------------------------------------
// Package problem_300
// Author: hanzhi
// Date: 2025/2/11
// -------------------------------------------------

package problem_300

import (
	"fmt"
	"math"
)

func lengthOfLISDP(nums []int) int {
	// 定义dp[i] 表示以i(包括i)为结尾的序列，一共有多长
	// 那么搜索的话也好说了
	// 只需要看n[i] > n[j]的时候，dp[j]有多长了，
	// dp[i] = dp[j] + 1
	if len(nums) == 1 {
		return 1
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dpValue := 0
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j] > dpValue {
					dpValue = dp[j]
				}
			}
		}
		dp[i] = dpValue + 1

	}
	maxValue := math.MinInt64
	for i := 0; i < len(nums); i++ {
		maxValue = max2(maxValue, dp[i])
	}
	fmt.Println(dp)
	return maxValue
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
