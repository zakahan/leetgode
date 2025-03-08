package problem_62

func UniquePaths(m int, n int) int {
	// 试一试动态规划把
	// dp[i,j]表示的是 到达[i,j]有几种方法

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := range dp {
		for j := range dp[0] {
			if i < len(dp)-1 {
				dp[i+1][j] += dp[i][j]
			}
			if j < len(dp[0])-1 {
				dp[i][j+1] += dp[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}
