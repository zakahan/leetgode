// -------------------------------------------------
// Package problem_221
// Author: hanzhi
// Date: 2024/12/24
// -------------------------------------------------

package problem_221

func maximalSquare(matrix [][]byte) int {
	// 动态规划解
	// dp 二维数组，表示以ij为右下角的矩形的边长，
	// dp[i,j] = 0 if mat[i,j] == 0
	// dp[i,j] = 1+ min(dp[i-1,j], dp[i,j-1], dp[i-1,j-1]) if mat[i,j] == 1
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	maxValue := 0
	// 开始插入
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 {
				if matrix[i][j] == '0' {
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			} else {
				if matrix[i][j] == '0' {
					dp[i][j] = 0
				} else {
					dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
			}
			if dp[i][j] > maxValue {
				maxValue = dp[i][j]
			}
		}
	}
	return maxValue * maxValue
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
