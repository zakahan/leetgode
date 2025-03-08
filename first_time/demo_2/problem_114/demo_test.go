package problem_114

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	var left = TreeNode{2, nil, nil}
	var root = TreeNode{1, &left, nil}
	flatten(&root)
	preOrder(&root)
}

func preOrder(node *TreeNode) {
	fmt.Print(" ")
	if node != nil {
		fmt.Print(node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	} else {
		fmt.Print("null")
	}
}
