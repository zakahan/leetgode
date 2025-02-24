// -------------------------------------------------
// Package problem_560
// Author: hanzhi
// Date: 2025/2/6
// -------------------------------------------------

package problem_560

/*
560. 和为 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。
注意是连续的！！！那就不能动态规划了
*/
func subarraySum(nums []int, k int) int {
	var res = 0
	myMap := make(map[int]int)
	myMap[0] = 1 // 什么都不选，前缀和为0的情况至少出现过一次
	// nums[i] 重新定义为累计sum了还是一个区间[i-j)  不包括自己
	var perixSum = 0
	for _, num := range nums {
		perixSum += num
		//fmt.Printf("当前前缀和为 %v, 当前num为 %v\n", perixSum, num)
		res += myMap[perixSum-k]
		// 更新哈希表
		myMap[perixSum] += 1
	}

	//fmt.Println(myMap)
	return res
}
