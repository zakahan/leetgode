// -------------------------------------------------
// Package problem_543
// Author: hanzhi
// Date: 2025/2/12
// -------------------------------------------------

package problem_543

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var MaxLength int

func diameterOfBinaryTree(root *TreeNode) int {
	// 这个路径全大于0，好做了
	// 计算当前节点最长边长，和当前节点为头的总长
	MaxLength = 1
	dfs(root)
	return MaxLength - 1

}

// 后根遍历
func dfs(root *TreeNode) int {
	if root != nil {
		left := dfs(root.Left)   // 求左子树的深度
		right := dfs(root.Right) // 求右子树的深度
		// 更新最大值
		if left+right+1 > MaxLength {
			MaxLength = left + right + 1
		}
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	} else {
		return 0
	}
}
