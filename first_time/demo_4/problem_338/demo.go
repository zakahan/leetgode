// -------------------------------------------------
// Package problem_338
// Author: hanzhi
// Date: 2025/1/31
// -------------------------------------------------

package problem_338

func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 { // 如果是偶数，那么就等于i/2之后的那个值
			res[i] = res[i/2]
		} else { // 奇数比前一个偶数多1
			res[i] = res[i-1] + 1
		}
	}
	return res
}
