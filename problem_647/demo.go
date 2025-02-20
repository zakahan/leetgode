// -------------------------------------------------
// Package problem_647
// Author: hanzhi
// Date: 2025/2/20
// -------------------------------------------------

package problem_647

func countSubstrings2(s string) int {
	ans := 0
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 1; i+j <= l; j++ {
			if isThat(s[i : i+j]) {
				//fmt.Println(s[i : i+j])
				ans++
			}
		}
	}
	return ans
}

func isThat(s string) bool {
	if len(s) == 0 {
		return false
	}

	l := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[l-i] {
			return false
		}
	}
	return true

}
