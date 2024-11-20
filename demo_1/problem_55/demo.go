package problem_55

// 时间复杂度n2
// e.g. : [2,3,1,1,4]
func canJumpOld(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		if dp[i] { // 能到达这里才行
			for k := 1; k <= nums[i]; k++ {
				if i+k == len(nums)-1 {
					return true
				}
				if i+k < len(nums) {
					dp[i+k] = true
				}
			}
		} else {
			//return false
		}
	}
	return dp[len(nums)-1]
}
