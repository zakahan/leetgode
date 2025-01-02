// -------------------------------------------------
// Package problem_236
// Author: hanzhi
// Date: 2025/1/2
// -------------------------------------------------

package problem_236

import (
	"fmt"
	"strconv"
)

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

var stackP []*TreeNode
var stackQ []*TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 用栈来解决，设计两个栈，分别存p q的祖宗节点
	stackP = make([]*TreeNode, 0)
	stackQ = make([]*TreeNode, 0)
	preOrder(root, p, &stackP)
	preOrder(root, q, &stackQ)
	//fmt.Println("啊哈")
	////fmt.Println(stackP)
	PrintNode(stackP)
	//fmt.Println("下一条！！！\n")
	PrintNode(stackQ)
	// 接下来判断祖宗
	var l int
	if len(stackP) < len(stackQ) {
		l = len(stackP)
	} else {
		l = len(stackQ)
	}
	for i := 0; i < l; i++ {
		if stackP[i] != stackQ[i] {
			return stackP[i-1]
		}
	}
	return stackP[l-1]
}
func preOrder(node, aim *TreeNode, stack *[]*TreeNode) bool {
	*stack = append(*stack, node)
	if node == nil {
		*stack = (*stack)[:len(*stack)-1]
		return false
	} else {
		// 是否命中？命中直接return true
		if node == aim {
			return true
		}
		// 没命中，但是还有希望
		if preOrder(node.Left, aim, stack) {
			// 表示未来命中了
			return true
		} else {
			// 这条线没命中，
		}
		if preOrder(node.Right, aim, stack) {
			return true
		} else {
			// 没命中
		}
		*stack = (*stack)[:len(*stack)-1]
		return false
	}
}

func PrintNode(s []*TreeNode) {
	for i := range s {
		if s[i] != nil {
			fmt.Printf(strconv.Itoa(s[i].Val) + " ")
		} else {
			fmt.Printf("nil")
		}

	}
	fmt.Println()
}
