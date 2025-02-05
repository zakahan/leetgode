// -------------------------------------------------
// Package problem_581
// Author: hanzhi
// Date: 2025/1/31
// -------------------------------------------------

package problem_581

func findUnsortedSubarray(nums []int) int {
	//初始化
	l := len(nums)
	min1 := nums[l-1]
	max1 := nums[0]
	var begin = 0
	var end = -1
	//遍历
	for i := 0; i < l; i++ {
		if nums[i] < max1 { //从左到右维持最大值，寻找右边界end
			end = i
		} else {
			max1 = nums[i]
		}

		if nums[l-i-1] > min1 { //从右到左维持最小值，寻找左边界begin
			begin = l - i - 1
		} else {
			min1 = nums[l-i-1]
		}
	}
	return end - begin + 1
}
