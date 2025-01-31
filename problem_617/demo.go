// -------------------------------------------------
// Package problem_617
// Author: hanzhi
// Date: 2025/1/31
// -------------------------------------------------

package problem_617

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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	root1 = preOrder(root1, root2)
	return root1

}

func preOrder(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = preOrder(root1.Left, root2.Left)
	root1.Right = preOrder(root1.Right, root2.Right)
	return root1
}
