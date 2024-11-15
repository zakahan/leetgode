package problem_55

// 第一个方法有问题，实际上只要找最远就可以了，如果你能到4，那么你一定能跳到3
func CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxCover := nums[0]
	// 如果能到达4 那么你一定能到3 所以 不用两个循环了，只需要一个循环，变量变成maxCover，如果挑不到了，自然就推出循环了

	for i := 0; i <= maxCover; i++ {
		if maxCover < i+nums[i] {
			maxCover = i + nums[i]
		}
		if maxCover >= len(nums)-1 {
			return true
		}
	}
	return false
}
