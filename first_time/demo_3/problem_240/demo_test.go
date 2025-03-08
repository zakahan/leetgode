// -------------------------------------------------
// Package problem_240
// Author: hanzhi
// Date: 2025/1/3
// -------------------------------------------------

package problem_240

import "testing"

func TestDemo(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30},
	}
	searchMatrix(matrix, 5)
}
