package problem_114

var pre *TreeNode = nil

func flatten(root *TreeNode) {
	pre = nil
	postOrder(root)
}

func postOrder(root *TreeNode) {
	// 后序遍历解
	if root == nil {
		return
	} else {
		postOrder(root.Right)
		postOrder(root.Left)
		root.Right = pre
		root.Left = nil
		pre = root
	}
}
