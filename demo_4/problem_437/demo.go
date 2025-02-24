// -------------------------------------------------
// Package problem_437
// Author: hanzhi
// Date: 2025/2/19
// -------------------------------------------------

package problem_437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var preSumDict map[int]int
var ans int

func pathSum(root *TreeNode, targetSum int) int {
	// 通过深度遍历来确定前缀和
	preSumDict = make(map[int]int) // 初始化，路径之和为0的，存在一个，就是空的那个
	preSumDict[0] = 1
	ans = 0
	dfs(root, 0, targetSum)
	return ans
}

func dfs(root *TreeNode, current, targetSum int) {
	if root == nil {
		return
	}
	// 查找是否存在前缀为current-targetSum的元素（就是当前值 + 前缀 = target
	current = current + root.Val
	ans += preSumDict[current-targetSum] // 如果存在，那么就多了一条线
	preSumDict[current]++
	dfs(root.Left, current, targetSum)
	dfs(root.Right, current, targetSum)
	preSumDict[current]-- // 结束，退出
}
