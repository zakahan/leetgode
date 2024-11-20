package problem_42

func Trap(height []int) int {
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

	// 找凸点
	var nowMaxIndex = 0
	// fixme:目前的思路类似快排 先找到一个最值点，然后找他左边的最指点和右边的最指点，递归
	for i := range height {
		if height[i] > height[nowMaxIndex] {
			nowMaxIndex = i
		}
	}
	outstandPoints[nowMaxIndex] = true
	if nowMaxIndex >= 1 {
		findMaxLeft(height, nowMaxIndex, &outstandPoints) // left只允许往left找！
	}
	if nowMaxIndex < len(outstandPoints)-1 {
		findMaxRight(height, nowMaxIndex, &outstandPoints) // right只允许往right找！
	}
	for i := range outstandPoints {
		if outstandPoints[i] && standI != -1 && standJ == -1 {
			standJ = i
		}
		if outstandPoints[i] && standI == -1 {
			standI = i
		}
	}
	if standJ == -1 {
		return 0
	}

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

func minH(h1 int, h2 int) int {
	if h1 < h2 {
		return h1
	} else {
		return h2
	}
}

func findMaxLeft(height []int, divided int, outstandPointsAd *[]bool) {

	if divided > 1 {
		//subHeight := height[:divided]
		var nowMaxIndex = 0 // 最好是从最左边开始
		for i := 0; i < divided; i++ {
			if height[i] > height[nowMaxIndex] {
				nowMaxIndex = i
			}
		}
		(*outstandPointsAd)[nowMaxIndex] = true

		if nowMaxIndex > 1 {
			findMaxLeft(height, nowMaxIndex, outstandPointsAd)
		} else {
			return
		}
	} else {
		return
	}

}

func findMaxRight(height []int, divided int, outstandPointsAd *[]bool) {

	if divided < len(height)-2 { // len(h)-d >2    d < len(h) -2
		//subHeight := height[divided+1:]
		var nowMaxIndex = len(height) - 1
		for i := divided + 1; i < len(height)-1; i++ {
			if height[i] > height[nowMaxIndex] {
				nowMaxIndex = i
			}
		}
		(*outstandPointsAd)[nowMaxIndex] = true
		if nowMaxIndex < len(height)-2 {
			findMaxRight(height, nowMaxIndex, outstandPointsAd)
		}
	} else {
		return
	}

}
