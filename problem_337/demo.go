// -------------------------------------------------
// Package problem_337
// Author: hanzhi
// Date: 2025/2/13
// -------------------------------------------------

package problem_337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var onItMap map[*TreeNode]int
var notOnItMap map[*TreeNode]int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	// 后序遍历
	dfs(root.Left)
	dfs(root.Right)
	// 计算
	// 1. 选当前节点，那么就不能选自己的两个子节点了
	onItMap[root] = notOnItMap[root.Right] + notOnItMap[root.Left] + root.Val
	// 2. 不选当前节点，那么就可以选自己的两个子节点了（但是也可以不选）
	notOnItMap[root] = maxA(onItMap[root.Right], notOnItMap[root.Right]) + maxA(onItMap[root.Left], notOnItMap[root.Left])
	return
}

func maxA(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(root *TreeNode) int {
	//
	onItMap = make(map[*TreeNode]int)
	notOnItMap = make(map[*TreeNode]int)
	dfs(root)
	if onItMap[root] > notOnItMap[root] {
		return onItMap[root]
	} else {
		return notOnItMap[root]
	}
}
