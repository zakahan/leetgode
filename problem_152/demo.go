package problem_152

import (
	"fmt"
)

// 动态规划，可以优化空间复杂度，但是算了就这样了。
func maxProduct(nums []int) int {
	// 状态转移方程，max{a_i, f_min(i-1) * a_i , f_max(i-1) * a_i}
	// 状态转移方程，min{a_i, f_min(i-1) * a_i , f_max(i-1) * a_i}
	if len(nums) == 1 {
		return nums[0]
	}
	max_value := nums[0]
	dp_max := make([]int, len(nums))
	dp_min := make([]int, len(nums))
	dp_max[0] = nums[0]
	dp_min[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp_max[i] = max3(nums[i], dp_max[i-1]*nums[i], dp_min[i-1]*nums[i])
		dp_min[i] = min3(nums[i], dp_max[i-1]*nums[i], dp_min[i-1]*nums[i])

		if dp_max[i] > max_value {
			max_value = dp_max[i]
		}
	}

	//fmt.Println(dp_max)
	fmt.Println(dp_min)
	return max_value
}

func min3(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func max3(a int, b int, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}
