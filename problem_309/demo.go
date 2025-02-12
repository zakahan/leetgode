// -------------------------------------------------
// Package problem_309
// Author: hanzhi
// Date: 2025/2/12
// -------------------------------------------------

package problem_309

func maxProfit(prices []int) int {
	// dp[i] 表示前i天内的最大收益
	// 状态转移 分为几部分 dp[i][j] = dp[j] + max(prices[i]-prices[k])
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3) // 因为j不能等于0，所以我直接把dp[i][0]用来表示i之前的累计最大值？
	}
	// 好几把难想到啊
	// dp[i][0] -> 当前持有一支股票
	// dp[i][1] -> 没有股票，也处于冷冻期中(啥也不能干）
	// dp[i][2] -> 没有股票，不处于冷冻期中

	dp[0][0] = -prices[0] // 持有一只股票，那就只能是本股了
	dp[0][1] = 0          // 这个是不存在的情况，
	dp[0][2] = 0          // 就是没有股票
	for i := 1; i < len(dp); i++ {
		// 持有一只股票;那么就是i-1的时候持有，现在也不管，或者是上一次就没有，现在买入了
		dp[i][0] = maxF(dp[i-1][0], dp[i-1][2]-prices[i])
		// 不持有股票，且处于冷冻期；那就代表是之前的那支股票，卖了
		dp[i][1] = dp[i-1][0] + prices[i]
		// 不持有股票，不处于冷冻期; 第一个情况就是i-1的时候不持有，现在还不持有，第二个情况是之前不持有，处于冷冻期，现在解冻
		dp[i][2] = maxF(dp[i-1][1], dp[i-1][2])
	}
	return maxF(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func maxF(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(prices)) // 因为j不能等于0，所以我直接把dp[i][0]用来表示i之前的累计最大值？
	}
	// 状态转移 dp[i][j] = 最大的前i-2次之和
	for i := 0; i < len(dp)-1; i++ {
		for j := i + 1; j < len(dp); j++ {
			dp[i][j] = dp[i-2][0] + (prices[j] - prices[i])
		}
	}
*/
