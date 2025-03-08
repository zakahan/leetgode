package problem_46

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 又是一个回溯问题
// 这次是排列 带有顺序的
// 应该是一个完整的完全二叉树

var resList = [][]int{}

func Permute(nums []int) [][]int {
	resList = [][]int{}
	targetList := []int{}
	backtracking(nums, targetList)
	return resList
}

func backtracking(nums []int, targetList []int) {
	if len(nums) == 0 {
		newTargetList := make([]int, len(targetList))
		copy(newTargetList, targetList)
		resList = append(resList, newTargetList)
		return
	} else {
		for i := 0; i < len(nums); i++ {
			newTargetList := make([]int, len(targetList))
			copy(newTargetList, targetList)
			newTargetList = append(newTargetList, nums[i])
			subNum := getNums(nums, i)
			backtracking(subNum, newTargetList)

		}
	}

}

func getNums(nums []int, j int) []int {

	subNums := make([]int, len(nums)-1)
	var i = 0
	var k = 0
	for i < len(nums) {
		if i == j {
			i++
		} else {
			subNums[k] = nums[i]
			k++
			i++
		}

	}
	return subNums
}
