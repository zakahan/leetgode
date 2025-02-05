// -------------------------------------------------
// Package problem_394
// Author: hanzhi
// Date: 2025/2/5
// -------------------------------------------------

package problem_394

func decodeString(s string) string {
	//output := ""
	var stackLetter []uint8
	var stackNumber []int
	for i := 0; i < len(s); i++ {
		item := s[i]
		// 3[a2[c]]
		// end <- a$ <- c
		// end <- 3 <- 2   -> cc -> acc -> acc acc acc
		if '1' <= item && item <= '9' {
			number := int(item - '0')
			for i+1 < len(s) {
				if s[i+1] == '[' {
					//那么一切ok，无所谓
					break
				} else if '0' <= s[i+1] && s[i+1] <= '9' {
					// 如果依然是数字，那么就累加
					number = number*10 + int(s[i+1]-'0')
					i++
				}
			}
			// 得看看后面还有没有数字
			// 入栈数字，等一会儿入栈字符
			stackNumber = append(stackNumber, number)
			// 循环，找Letter字符串
		} else if item == ']' {
			// 出栈
			var j int = len(stackLetter) - 1
			for ; stackLetter[j] != '[' && j >= 0; j-- {
				// 遍历，直到找到了字符串起点
			}
			var loopString string = string(stackLetter[j+1:])
			// 字符出栈
			stackLetter = stackLetter[:j]
			// 数字出栈
			p := stackNumber[len(stackNumber)-1]
			stackNumber = stackNumber[0 : len(stackNumber)-1]
			for k := 0; k < p; k++ {
				// loopString的部分应该继续入栈，入字符栈
				for l := 0; l < len(loopString); l++ {
					stackLetter = append(stackLetter, loopString[l])
				}
			}
		} else {
			// 如果是左括号或者是字母，则入栈
			stackLetter = append(stackLetter, item)
		}

	}

	return string(stackLetter)
}
