// -------------------------------------------------
// Package problem_73
// Author: hanzhi
// Date: 2025/2/26
// -------------------------------------------------

package problem_73

import "math"

// 偷懒法，直接利用go的优势
func setZeroes(matrix [][]int) {
	value := math.MinInt64
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				for ik := 0; ik < len(matrix); ik++ {
					if matrix[ik][j] != 0 {
						matrix[ik][j] = value
					}
				}
				for jk := 0; jk < len(matrix[i]); jk++ {
					if matrix[i][jk] != 0 {
						matrix[i][jk] = value
					}

				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == value {
				matrix[i][j] = 0
			}
		}
	}

}
