// -------------------------------------------------
// Package problem_461
// Author: hanzhi
// Date: 2025/1/31
// -------------------------------------------------

package problem_461

func hammingDistance(x int, y int) int {
	z := x ^ y
	k := 0
	for z > 0 {
		k += z % 2
		z = z / 2
	}
	return k
}
