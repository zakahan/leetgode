package problem_84

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, n)
	maxArea := 0
	top := -1
	for i := 0; i < n; i++ {
		for top > -1 && heights[i] < heights[stack[top]] {
			for top > 0 && heights[stack[top]] > heights[i] { // 必须大于，否则不能出栈
				left := stack[top-1]
				w := i - left - 1
				h := heights[stack[top]]
				tmpArea := w * h
				if tmpArea > maxArea {
					maxArea = tmpArea
				}
				top--
			}
			if top == 0 && heights[stack[top]] > heights[i] {
				w := i - (-1) - 1 // 表示他更前面的也是   其实就
				h := heights[stack[top]]
				tmpArea := w * h
				if tmpArea > maxArea {
					maxArea = tmpArea
				}
				top--
			}
		}
		// 正常递增的话，不用管
		top++
		stack[top] = i
	}
	// 这个时候 栈内可能还有数据  一定是栈顶的大于下面的！
	for top > 0 { // 必须大于，否则不能出栈  (至少有两个
		left := stack[top-1]
		w := n - left - 1
		h := heights[stack[top]]
		tmpArea := w * h
		if tmpArea > maxArea {
			maxArea = tmpArea
		}
		top--
	}
	if top == 0 {
		// 就剩一个了，这个就是最矮的
		w := n
		h := heights[stack[top]]
		tmpArea := w * h
		if tmpArea > maxArea {
			maxArea = tmpArea
		}
		top--
	}
	return maxArea
}
