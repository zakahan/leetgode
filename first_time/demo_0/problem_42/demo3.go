package problem_42

// 官方解1：动态规划法

func trapDP(height []int) int {
	leftMaxList := make([]int, len(height))
	rightMaxList := make([]int, len(height))
	leftMaxValue := 0
	rightMaxValue := 0

	for i := range leftMaxList {
		if height[i] > leftMaxValue {
			leftMaxValue = height[i]
		}
		leftMaxList[i] = leftMaxValue
	}

	for j := len(height) - 1; j >= 0; j-- {
		if height[j] > rightMaxValue {
			rightMaxValue = height[j]
		}
		rightMaxList[j] = rightMaxValue
	}
	// 最终用leftMaxList作为能接到的雨水的表
	res := 0
	for i := range leftMaxList {
		leftMaxList[i] = minInt(leftMaxList[i], rightMaxList[i])
		//fmt.Print(leftMaxList[i])
		res += leftMaxList[i] - height[i]
	}
	return res
}

func minInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
