package problem_121

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/
func maxProfit1(prices []int) int {
	// 看着就是dp把？
	// dp[i] 表示在0-i-1期间买入， 在i天卖出的价格
	minPrices := make([]int, len(prices)) // 表示0~（i-1）期间的最小值
	minValue := 100000 + 1
	for i := range prices {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		minPrices[i] = minValue
	}
	// 其次
	dp := make([]int, len(prices))
	dp[0] = 100000 + 1
	if len(prices) <= 1 {
		return 0
	}
	var maxProfitValue = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = prices[i] - minPrices[i-1]
		if dp[i] < 0 {
			dp[i] = 0
		}
		if maxProfitValue < dp[i] {
			maxProfitValue = dp[i]
		}
	}
	return maxProfitValue
}
