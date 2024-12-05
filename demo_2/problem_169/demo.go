package problem_169

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

// 简单的 哈希表法 就不写了

func majorityElement(nums []int) int {
	key := nums[0]
	value := 1
	for i := 1; i < len(nums); i++ {
		// 如果当前值等于key 那么value++
		// 如果不等于key 那么value--
		// 一旦value小于0了，下次那么就更新key
		if nums[i] == key {
			value++
		} else {
			if value == 0 {
				key = nums[i]
				value = 1
			} else {
				value--
			}
		}
	}
	return key
}
