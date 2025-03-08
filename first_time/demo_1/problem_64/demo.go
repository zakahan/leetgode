package problem_64

func MinPathSum(grid [][]int) int {
	// 还是动态规划
	// 因为上面的不会受到下面的影响，所以从上往下，遍历是没问题的
	// 内循环从右往左也没问题，因为内外不会相互影响的
	// dp表示到达当前位置所需要的最小的代价
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		for j := range dp[0] {
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] = minInt(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if j-1 >= 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i-1 >= 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = grid[i][j]
			}

		}
	}
	return dp[m-1][n-1]
}

func minInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
