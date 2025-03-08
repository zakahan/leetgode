// -------------------------------------------------
// Package problem_239
// Author: hanzhi
// Date: 2025/1/3
// -------------------------------------------------

package problem_239

// 这个题有点像之前那个最小栈。
// 能不能做两个队列，实现一个常规队列 + 一个“最大值队列”
// 最大值队列是单调递减的！！
func maxSlidingWindow(nums []int, k int) []int {

	var maxQueue []int
	var normalQueue []int
	var result []int
	var l = len(nums)

	normalQueue = append(normalQueue, nums[0])
	maxQueue = append(maxQueue, nums[0])

	// 初始化队列
	for i := 1; i < k; i++ {
		// 常规队列就正常插入
		normalQueue = append(normalQueue, nums[i])
		if nums[i] <= maxQueue[len(maxQueue)-1] {
			// 如果小于等于最大值的队尾，那就正常插入
			maxQueue = append(maxQueue, nums[i])
		} else {
			// 如果这个值比 队尾大，那就要去除所有比他小的值
			for j := len(maxQueue) - 1; j >= 0 && maxQueue[j] < nums[i]; j-- {
				maxQueue = maxQueue[:j]
			}
			maxQueue = append(maxQueue, nums[i])
		}
	}

	//fmt.Println("初始情况")
	//fmt.Println(maxQueue)

	// 此时可以把最大值加进去
	result = append(result, maxQueue[0])
	for i := k; i < l; i++ {
		// 先删除
		item := normalQueue[0]
		normalQueue = normalQueue[1:]
		// 判断item和最大值的比较，如果是同一个，那么最大值要去除
		if item == maxQueue[0] {
			maxQueue = maxQueue[1:]
		}

		// 常规队列就正常插入
		normalQueue = append(normalQueue, nums[i])
		if len(maxQueue) == 0 || nums[i] <= maxQueue[len(maxQueue)-1] {
			// 如果小于等于最大值的队尾，那就正常插入
			maxQueue = append(maxQueue, nums[i])
		} else {
			// 如果这个值比 队尾大，那就要去除所有比他小的值
			for j := len(maxQueue) - 1; j >= 0 && maxQueue[j] < nums[i]; j-- {
				maxQueue = maxQueue[:j]
			}
			maxQueue = append(maxQueue, nums[i])
		}

		// 打印
		//fmt.Printf("第%v次情况\n", i)
		//fmt.Println(maxQueue)
		result = append(result, maxQueue[0])
	}
	return result
}
