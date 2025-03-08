// -------------------------------------------------
// Package second_time
// Author: hanzhi
// Date: 2025/3/5
// -------------------------------------------------

package second_time

func twoSum(nums []int, target int) []int {

	// 考察点：哈希表
	myMap := make(map[int]int)

	for i, num := range nums {
		if myMap[target-num] != 0 {
			return []int{i, myMap[target-num] - 1}
		} else {
			myMap[num] = i + 1
		}
	}
	return []int{}

}
