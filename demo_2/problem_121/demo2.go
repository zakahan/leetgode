package problem_121

func maxProfit(prices []int) int {
	// 看着就是dp把？
	// dp[i] 表示在0-i-1期间买入， 在i天卖出的价格
	minValue := prices[0]
	maxProfitValue := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		if prices[i]-minValue > maxProfitValue {
			maxProfitValue = prices[i] - minValue
		}
	}
	return maxProfitValue
}
