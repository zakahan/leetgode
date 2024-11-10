package problem_32

type Item struct {
	Val   rune
	Index int
}

func LongestValidParentheses(s string) int {
	s = ")" + s
	n := len(s)

	if n <= 1 {
		return 0
	}
	if n == 2 {
		if s == "()" {
			return 2
		} else {
			return 0
		}
	}

	recalls := make([]int, n)
	var stack []Item = []Item{}
	for i, value := range s {
		if value == ')' {
			if len(stack) > 0 && stack[len(stack)-1].Val == '(' {
				stack = stack[0 : len(stack)-1]
				//fmt.Printf("%s号元素进栈，内容%s\n", strconv.Itoa(i), string(value))
			} else { // 先出现，或者没有对应的匹配点代表错误，记录错误点
				recalls[i] = 1
				stack = []Item{} // 清空栈
			}

		} else if value == '(' {
			stack = append(stack, Item{'(', i})
			recalls[i] = 0
		}
	}
	if len(stack) != 0 {
		for k := range stack {
			index := stack[k].Index
			recalls[index] = 1
		}
	}

	// 也就是说，这些1之间的就是有效的子串了
	var i = 0
	var j = 0
	maxLength := j - i
	for i < len(recalls) && j < len(recalls) {
		if recalls[i] == 1 && recalls[j] == 1 {
			if j-i-1 > maxLength {
				maxLength = j - i - 1
			}
			i = j
		}
		j++

	}

	if j-i-1 > maxLength {
		maxLength = j - i - 1
	}
	return maxLength
}
