// -------------------------------------------------
// Package problem_226
// Author: hanzhi
// Date: 2025/1/2
// -------------------------------------------------

package problem_226

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

func invertTree(root *TreeNode) *TreeNode {
	ReverseLR(root)
	return root
}

func ReverseLR(node *TreeNode) {
	if node == nil {
		return
	} else {
		// 不用管是否为nil，直接交换
		node.Left, node.Right = node.Right, node.Left
		ReverseLR(node.Right)
		ReverseLR(node.Left)
	}
}
