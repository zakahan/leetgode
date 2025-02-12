// -------------------------------------------------
// Package problem_322
// Author: hanzhi
// Date: 2025/2/12
// -------------------------------------------------

package problem_322

import "math"

func coinChange(coins []int, amount int) int {
	// 考虑用动态规划来解决
	// dp[i] 表示当值为i的时候，最少通过多少个硬币来达成目标
	dp := make([]int, amount+1)
	// dp[i] = dp[i-conins[j]] + 1
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				if dp[i-coins[j]] != -1 && dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
		if dp[i] == math.MaxInt {
			dp[i] = -1
		}
	}
	return dp[amount]
}
