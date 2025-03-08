// -------------------------------------------------
// Package problem_301
// Author: hanzhi
// Date: 2025/2/11
// -------------------------------------------------

package problem_301

import "fmt"

func isValid(s string) bool {
	x := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			x++
		}
		if s[i] == ')' {
			x--
			if x < 0 {
				return false
			}
		}
	}
	if x == 0 {
		return true
	} else {
		return false
	}
}

func backcall(ans *[]string, s string, start, lRemove, rRemove int) {
	fmt.Println(s)
	if lRemove == 0 && rRemove == 0 {
		if isValid(s) {
			*ans = append(*ans, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] {
			// 表示 (((((的情况直接跳过
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lRemove+rRemove > len(s)-i {
			return
		}
		// 尝试抹去左括号
		if lRemove > 0 && s[i] == '(' {
			backcall(ans, s[:i]+s[i+1:], i, lRemove-1, rRemove)
		}
		// 尝试抹去右括号
		if rRemove > 0 && s[i] == ')' {
			backcall(ans, s[:i]+s[i+1:], i, lRemove, rRemove-1)
		}
	}
}

func removeInvalidParentheses(s string) []string {
	lRemove, rRemove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lRemove++
		} else if ch == ')' {
			if lRemove == 0 {
				rRemove++
			} else {
				lRemove--
			}
		}
	}

	res := []string{}
	// 开始回溯
	backcall(&res, s, 0, lRemove, rRemove)
	return res
}
