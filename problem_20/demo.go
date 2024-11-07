package problem_20

func IsValid(s string) bool {
	// 这把用栈
	stack := make([]rune, 10000)
	var top = -1
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			// 入栈
			top++
			stack[top] = c
		}
		if c == ')' || c == ']' || c == '}' {
			// 取栈顶 判断是否匹配
			if top < 0 {
				return false
			}
			topVar := stack[top]
			if !isPaired(topVar, c) {
				return false
			}
			top-- // 出栈

		}
	}
	if top == -1 {
		return true
	} else {
		return false
	}
}

func isPaired(char1 rune, char2 rune) bool {
	if char1 == '(' && char2 == ')' {
		return true
	}
	if char1 == '[' && char2 == ']' {
		return true
	}
	if char1 == '{' && char2 == '}' {
		return true
	}
	return false
}
