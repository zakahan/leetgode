package problem_78

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的
子集
（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

还是回溯方案吧
*/
var resList [][]int

func Subsets(nums []int) [][]int {
	resList = [][]int{}
	resList = append(resList, []int{})
	backtracking(len(nums), 0, nums, []int{})
	return resList
}

func backtracking(depth int, start int, nums []int, list []int) {
	if depth == 0 {
		return
	} else {
		for i := start; i < len(nums); i++ {
			newList := make([]int, len(list), len(list)+1)
			copy(newList, list)
			newList = append(newList, nums[i])
			resList = append(resList, newList)
			backtracking(depth-1, i+1, nums, newList)
		}
	}
}
