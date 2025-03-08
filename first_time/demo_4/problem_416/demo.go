package problem_416

import "fmt"

// 这个问题是如何转化为01背包的
// 1. 假设nums总和为x，x如果是奇数一定不行，x如果是偶数，那么就意味着是从Nums里面选取k个，使得这k个数据的和是x/2
func CanPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}
	// dp[i,j] 前0~i个， 总和为j
	dp := make([][]bool, len(nums)+1)
	for k := 0; k < len(nums)+1; k++ {
		dp[k] = make([]bool, target+1)
	}
	// 第0行 和第0列什么用都没有
	dp[0][nums[0]] = true // 起点，第一个放第一个里有用
	for i := 1; i < len(nums); i++ {
		// i表示的是前0~i个 i是行
		for j := 0; j <= target; j++ {
			// 第一种是刚好补上缺口，也就是在前i-1个里，总数j = nums[i]（新来的） + old
			// 保证old存在
			old := j - nums[i]
			// 第二种是前i-1个里，就已经满足总数是j了，那么就不用考虑新的了
			// 所以其实就两种选项
			if old >= 0 {
				dp[i][j] = dp[i-1][old] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[len(nums)-1][target]
}
