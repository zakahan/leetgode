// -------------------------------------------------
// Package problem_124
// Author: hanzhi
// Date: 2025/2/8
// -------------------------------------------------

package problem_124

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func maxPathSum(root *TreeNode) int {
	// 深度优先遍历
	maxSum = math.MinInt64
	dfs(root)
	return maxSum
}

func dfs(root *TreeNode) int {
	if root != nil {
		// 左子树的最大贡献值

		left := dfs(root.Left)
		right := dfs(root.Right)
		selfMaxP := 0

		if left < 0 {
			left = 0
		}
		if right < 0 {
			right = 0
		}
		selfMaxSum := root.Val + left + right

		// 看看是否大于0
		if left > right {
			selfMaxP = root.Val + left
		} else {
			selfMaxP = root.Val + right // 返回当前节点的最大路径之
		}
		if selfMaxSum > maxSum {
			maxSum = selfMaxSum
		}
		//fmt.Printf("当前节点 %v，最大贡献为 %v，当前maxSum为%v\n", root.Val, selfMaxP, maxSum)
		return selfMaxP // 给的是只有一条线+自己的， 对比的时候是两条线+自己的
	} else {
		return 0
	}
}
