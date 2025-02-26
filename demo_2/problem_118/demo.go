// -------------------------------------------------
// Package problem_118
// Author: hanzhi
// Date: 2025/2/26
// -------------------------------------------------

package problem_118

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = res[i-1][j] + res[i-1][j-1]
			}
			res[i] = append(res[i], value)
		}
	}
	return res

}
