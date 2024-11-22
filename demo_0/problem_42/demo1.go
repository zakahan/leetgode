package problem_42

// 官方题解2 单调栈法

func trap(height []int) int {
	res := 0
	// 定义单调栈
	stack := make([]int, len(height))
	for i := range stack {
		stack[i] = -1
	}
	top := -1

	//
	for i := range height {
		// 只要单调递减的入栈，递增了，出栈，
		for top > -1 && height[i] > height[stack[top]] {
			oldTopValue := height[stack[top]]
			nowValue := height[i]
			stack[top] = -1
			top--
			if top == -1 {

				break
			}
			leftValue := height[stack[top]]
			// 计算宽度
			w := i - stack[top] - 1
			h := minHInt(nowValue, leftValue) - oldTopValue
			//stack[top] = -1
			//top--
			res += w * h

		}
		// 无论是递增还是递减，都会入栈
		top++
		stack[top] = i

	}

	// 如果此时站内还有元素呢？那接不到水啊
	return res
}

func minHInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
