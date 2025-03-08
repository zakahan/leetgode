// -------------------------------------------------
// Package problem_238
// Author: hanzhi
// Date: 2025/1/3
// -------------------------------------------------

package problem_238

// 求两个东西，一个是前缀积一个是后缀积

func productExceptSelf(nums []int) []int {
	preHrix := make([]int, len(nums))
	subHrix := make([]int, len(nums))
	preHrix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preHrix[i] = preHrix[i-1] * nums[i]
	}

	subHrix[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		subHrix[i] = subHrix[i+1] * nums[i]
	}

	//fmt.Println(preHrix)
	//fmt.Println(subHrix)
	res := make([]int, len(nums))

	res[0] = subHrix[1]
	res[len(nums)-1] = preHrix[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		res[i] = preHrix[i-1] * subHrix[i+1]

	}
	return res
}
