package problem_104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return getDepth(root, 1)
}

func getDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth - 1 // 就是不算自己这一层了
	} else {
		return maxInt(getDepth(node.Left, depth+1), getDepth(node.Right, depth+1))
	}
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
