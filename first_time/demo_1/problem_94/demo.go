package problem_94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var resList []int

func inorderTraversal(root *TreeNode) []int {
	// 中序遍历
	resList = []int{}
	Order(root)
	return resList
}

func Order(node *TreeNode) {
	if node != nil {
		Order(node.Left)
		x := node.Val
		resList = append(resList, x)
		Order(node.Right)
	} else {
		return
	}
}
