// -------------------------------------------------
// Package problem_621
// Author: hanzhi
// Date: 2025/2/5
// -------------------------------------------------

package problem_621

func leastInterval(tasks []byte, n int) int {
	// 就是模拟一个单进程CPU的任务调度
	var output = 0
	var taskNumber = len(tasks) // 表达有多少个任务
	state := make([]int, 26)
	busy := make([]int, 26) // 表示当前哪个进程繁忙中（使用了之后，就倒计时记录为2，每次刷新更新）
	for _, b := range tasks {
		state[int(b-'A')] += 1
	}

	for taskNumber > 0 {
		// 先更新一下
		for key := range busy {
			if busy[key] > 0 {
				busy[key]-- // 表示更新一下
			}
		}

		// 看一下谁没任务，那就谁来
		var flag = 0
		var execTaskKey = 27 // 表示需要执行的任务ID，第一次的话默认为27，不会出现
		for key := range state {
			// 这里应该有个贪心的点，就是哪个任务最多，让谁先来，
			if state[key] != 0 && busy[key] == 0 { // 代表当前key没有全执行完，且并不繁忙
				if execTaskKey == 27 || state[key] > state[execTaskKey] {
					// 如果是第一次，赋值； 如果当前的值更大，那么也赋值
					execTaskKey = key
				}
				flag = 1

				//break // 退出for循环
			}
		}
		if flag == 0 {
			//表示都很忙，那么就空下来
			//fmt.Printf("没有任务可以执行，剩余任务数量 %v\n", taskNumber)
			output += 1
		} else {
			// 表示执行
			//fmt.Printf("执行任务 %v 剩余任务数量 %v 任务剩余个数 %v\n", string(execTaskKey), taskNumber)
			busy[execTaskKey] = n + 1 // 表示还有n轮空窗期
			output += 1               // 当前key占用cpu
			taskNumber--              // 表示当前task消失
			state[execTaskKey]--
		}

	}

	return output
}
