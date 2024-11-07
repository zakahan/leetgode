package problem_42

// Error： 这是个错误方法，先留着，我之后改进。
// 接雨水，经典难题，看看今天下午，这一个下午能不能搞定。
//搞不懂我就看答案呗，能咋办（我好像可以做出来！！）
//
/*
*
第P[i,j]接到的雨水量由什么决定？
左侧极值点和右侧极值点 他们两个的最小值决定
首先提取突点， 其意义为 h[i] > h[i+1] and h[i] > h[i-1]

*
*/
func TrapError(height []int) int {
	var totalWater = 0
	// 获取凸点矩阵
	//var outstandPoints []bool
	outstandPoints := make([]bool, len(height))
	// 先排除特殊条件小于2，一点接受不到,至少三个
	if len(height) <= 2 {
		return 0
	}

	var standI = -1
	var standJ = -1
	for i := range height {
		if i == 0 {
			if height[i] > height[i+1] {
				// 前者比后者大才行
				outstandPoints[i] = true
			} else {
				outstandPoints[i] = false
			}
		} else if i == len(height)-1 {
			if height[i] > height[i-1] {
				// 他比后者大
				outstandPoints[i] = true
			} else {
				outstandPoints[i] = false
			}
		} else {
			// 常规情况
			if height[i] > height[i+1] && height[i] > height[i-1] {
				outstandPoints[i] = true
			}
		}
		// 找起始两个凸点
		if outstandPoints[i] && standI != -1 && standJ == -1 {
			standJ = i
		}
		if outstandPoints[i] && standI == -1 {
			// 是凸点，且之前stand_i 没加过
			standI = i
		}

	} // 结束凸点的记录

	// 凸点和凸点之间才能接到雨水
	// 继续排除特殊情况：如果i和j有一个还是-1（代表只有一个极值点）
	if standI == -1 || standJ == -1 {
		return 0
	}
	// fixme 接雨水的这部分是没错的，错的是找“极值点”的方法
	// 或者说找极值点本身是错的
	var m = standJ
	for m < len(outstandPoints) {
		// 开始接雨水
		var heightLine = minH(height[standI], height[standJ])
		for i := standI; i < standJ; i++ { // 这个范围没错，在standJ之前
			if heightLine > height[i] {
				totalWater += heightLine - height[i]
			} else {
				totalWater += 0
			}
		}
		// 当一次遍历结束，游动指针
		standI = standJ

		for m = standJ + 1; m < len(outstandPoints); m++ {
			if outstandPoints[m] {
				standJ = m
				break

			}
		} // 当循环结束代表讲进入下一轮的接雨水了

	}
	return totalWater

}

func minHError(h1 int, h2 int) int {
	if h1 < h2 {
		return h1
	} else {
		return h2
	}
}
