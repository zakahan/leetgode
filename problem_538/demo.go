// -------------------------------------------------
// Package problem_538
// Author: hanzhi
// Date: 2025/2/6
// -------------------------------------------------

package problem_538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 很明显的后序遍历
func convertBST(root *TreeNode) *TreeNode {
	proOrder(root, 0)
	return root
}

func proOrder(node *TreeNode, sum int) int {
	if node != nil {
		sum = proOrder(node.Right, sum)
		//fmt.Printf("当前node %v | 当前sum %v \n", node.Val, sum)
		node.Val += sum
		sum = proOrder(node.Left, node.Val)
		return sum
	} else {
		return sum
	}
}
