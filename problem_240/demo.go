// -------------------------------------------------
// Package problem_240
// Author: hanzhi
// Date: 2025/1/3
// -------------------------------------------------

package problem_240

func searchMatrix(matrix [][]int, target int) bool {
	// 看题解了，大概就是 从 i=0 j = m-1的位置开始
	// 类似二叉搜索树， 如果比node小，那么就i--,大，就j++
	// i行 j列， m行n列
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1
	node := matrix[i][j]
	for i < m && j < n && i >= 0 && j >= 0 {
		node = matrix[i][j]
		//fmt.Println(node)
		if target == node {
			return true
		} else if target < node {
			j--
		} else {
			i++
		}

	}
	return false
}
