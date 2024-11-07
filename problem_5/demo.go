package problem_5

func LongestPalindrome(s string) string {
	// 这次直接动态规划来吧

	n := len(s)

	if n <= 1 {
		return s
	}

	matrix := make([][]bool, n)
	for i, _ := range matrix {
		matrix[i] = make([]bool, n)
	}
	// matrix表示状态转移矩阵，也就是P[i,j] 意思是从i到j的子串是否是回文
	// P[i,j] = P[i+1, j-1] && s[i]==s[j] （条件为0<i,j<n)
	// 边界条件， 单个字符肯定没问题，无字符这个表示不了  只需要矩阵的上三角部分

	for i := range matrix {
		matrix[i][i] = true
	}
	max_width_string := s[0:1]
	// 问题是咋扩展
	// 宽度为2的可能得我挨个实验了，淦
	for i := range matrix {
		if i-1 >= 0 {
			matrix[i-1][i] = s[i] == s[i-1]
			//
			if matrix[i-1][i] {
				max_width_string = s[i-1 : i+1]
			}
		}
		if i+1 < n {
			matrix[i][i+1] = s[i] == s[i+1]
			//
			if matrix[i][i+1] {
				max_width_string = s[i : i+2]
				print()
			}
		}
	}

	// 然后应该就可以迭代了,应该从宽度为0开始
	for width := 2; width < n; width++ { // 实际上width +1 才是看宽度的值
		var flag = false
		for i := range matrix { // i是起点
			if width+i < n {
				matrix[i][i+width] = matrix[i+1][i+width-1] && s[i] == s[i+width]
				if matrix[i][i+width] && width+1 > len(max_width_string) {
					// 如果这个是true 而且width要比原有的width大，那么就
					//fmt.Println(s[i : i+width+1])
					max_width_string = s[i : i+width+1]
					flag = true
				}
			}
		}
		if flag == false && width > len(max_width_string)+1 { //  +1 = -1 + 2 现在的计算宽度已经比最大历史宽度多两位了，
			break
		}
		flag = false
		// 如果宽度增长超过了最大历史记录两个点了还没有增长的话，那么就可以放弃了

	}

	// 遍历完毕，取寻找true在哪
	return max_width_string
}
