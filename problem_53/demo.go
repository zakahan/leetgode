package problem_53

// 哎，我有个好主意
func MaxSubArray(nums []int) int {
	var totalValue = -10001
	if len(nums) == 0 {
		return nums[0]
	}

	for i := range nums {
		if nums[i] > totalValue {
			totalValue = nums[i]
		}
	}
	if totalValue <= 0 {
		return totalValue
	}
	// 定义dp[i] i表示0~i个数据的情况下能拿到多少value(连续的）
	dp := make([]int, len(nums))

	var j = 0
	if nums[0] > 0 {
		j++
		dp[0] = nums[0]
	} else {
		dp[0] = 0
		j++
	}

	for j < len(nums) {
		dp[j] = dp[j-1] + nums[j] // 尝试加上这个值
		if dp[j] > 0 {
			// 如果加上之后还是正数 ，那就继续
			if dp[j] > totalValue {
				totalValue = dp[j]
			}
			j++
		} else {
			// 如果变成了复数，那么拜拜
			dp[j] = 0
			j = j + 1
		}
	}

	return totalValue
}
