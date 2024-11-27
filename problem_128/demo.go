package problem_128

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 时间O(n) 那我就只能勉为其难的来一波哈希表了
*/
func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	myMap := make(map[int]int)
	for i := range nums {
		myMap[nums[i]] = 1 // 为1表示没有遍历过
	}
	// 然后就是遍历map了
	var longestC = 0

	for _, value := range nums {
		if myMap[value] == 1 {
			// 接下来应该往左右查询
			myMap[value] = 2 // 2表示遍历过了
			left := value - 1
			right := value + 1
			for myMap[left] == 1 { //  往左一直走
				myMap[left] = 2 // 表示修改过了，防止以后再遍历它
				left--
			}
			for myMap[right] == 1 {
				myMap[right] = 2
				right++
			}

			tmp := (right - 1) - (left + 1) + 1
			if tmp > longestC {
				longestC = tmp
			}

		}
	}
	return longestC
}
