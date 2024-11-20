package problem_11

func maxArea(height []int) int {
	// 双指针法
	i := 0
	j := len(height) - 1
	var area int = 0
	for i < j {
		tmp := minFunc(height[i], height[j]) * (j - i)
		if tmp > area {
			area = tmp
		}

		// 开始双指针变化
		if height[i] < height[j] {
			// i小 移动i
			i = i + 1
		} else {
			j = j - 1
		}

	}
	return area
}

func minFunc(h1 int, h2 int) int {
	if h1 < h2 {
		return h1
	} else {
		return h2
	}
}
