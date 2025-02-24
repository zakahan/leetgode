// -------------------------------------------------
// Package problem_494
// Author: hanzhi
// Date: 2025/2/5
// -------------------------------------------------

package problem_494

// 这应该是一个01背包问题

func findTargetSumWays(nums []int, target int) int {

	// 先求最大值和最小值，都是有可能的
	// 带个偏移量吧
	var sumNums = 0
	for _, k := range nums {
		sumNums += k
	}

	if target > sumNums || target < -sumNums { // 如果目标值过大，比总和还大，那么一定是不存在的，那么就直接return 0
		return 0
	}

	var dp [][]int
	dp = make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2*sumNums+1)
	}
	// dp[i][j]表示的是 前0~i个数组元素中，运算和为j的情况
	//条件转移  dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j+nums[i]  (前者是 当前值取 nums[j],后者是当前值取-nums[j]
	b := sumNums // b表示偏置量
	//fmt.Printf("偏移量 : %v\n", b)
	//fmt.Println(len(dp))
	//fmt.Println(len(dp[0]))
	dp[0][nums[0]+b] += 1 // 必须是 +=而不是= 考虑一个特殊情况，就是[0] -> dp[0][0+b]和dp[-0+b]其实是一个位置
	dp[0][-nums[0]+b] += 1

	// 此外j方向要加上偏移量
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ { // j实际上已经带上偏置量了，j := trueJ + b, 比如trueJ=0 -> j =sumNums
			// j 表示和为j
			if j-nums[i] >= 0 && j-nums[i] < len(dp[i]) {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
			if j+nums[i] >= 0 && j+nums[i] < len(dp[i]) {
				dp[i][j] += dp[i-1][j+nums[i]]
			}
		}
	}

	return dp[len(nums)-1][target+b]
}
