package problem_152

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	// [1,2,-2,5,6,7,0,10]
	// 牢记，动态规划是，我只需要知道你能到这里，但是我不需要知道你怎么到这里的，
	// 第一种情况，正数和负数都考虑，就是不考虑0
	if len(nums) == 1 {
		return nums[0]
	}
	x := dpNormal(nums)
	y := dpAbnormal(nums)
	if x > y {
		return x
	} else {
		return y
	}
}

func dpNormal(nums []int) int {
	maxValue := math.MinInt32
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	i := 1
	for i < len(nums) {
		if dp[i-1] != 0 {
			dp[i] = dp[i-1] * nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > maxValue {
			maxValue = dp[i]
		}
		i++
	}
	fmt.Println("nor", dp)
	return maxValue
}

func dpAbnormal(nums []int) int {
	maxValue := math.MinInt32
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	i := 1
	for i < len(nums) {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] * nums[i]
		}
		if dp[i] > maxValue {
			maxValue = dp[i]
		}
		i++
	}
	fmt.Println("abn", dp)

	return maxValue
}
