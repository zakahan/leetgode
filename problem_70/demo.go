package problem_70

func climbStairs(n int) int {
	// 状态转移方程
	// 对于到达第i层，有P[i]中方式，
	// P[i,j] = P[i-1] + P[i-2]		// 如果他们存在的话

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 0 // 无所谓了
	dp[1] = 1
	dp[2] = 2

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
