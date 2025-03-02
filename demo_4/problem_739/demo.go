// -------------------------------------------------
// Package problem_739
// Author: hanzhi
// Date: 2025/2/19
// -------------------------------------------------

package problem_739

func dailyTemperatures(temperatures []int) []int {
	dayTemp := make([]int, len(temperatures))

	stack := make([]int, 0)
	for i, temp := range temperatures {

		last := len(stack) - 1
		for len(stack) != 0 && temperatures[stack[last]] < temp {
			// 如果当前栈内记录的温度，小于当前遍历到的温度，那么就修改内容了
			//fmt.Printf("开始出栈 %v\n", last)
			dayTemp[stack[last]] = i - stack[last]
			// 然后出栈
			stack = stack[:last]
			// 重新标记last
			last = len(stack) - 1
		}
		last = len(stack) - 1
		// 看一看当前元素 i 是否能入栈
		if len(stack) == 0 || temperatures[stack[last]] >= temp {
			//fmt.Printf("入栈 %v\n", i)
			stack = append(stack, i)
		}
	}
	return dayTemp
}
