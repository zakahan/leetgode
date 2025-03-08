package problem_114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten1(root *TreeNode) {
	// 怎么做原地操作？  移花接木解
	// 分为三个部分，根 左， 右
	// 将其变为 根  然后原本左边的挂载到右边， 原本右边的挂载到左边的最右侧  不断的套娃这个过程
	if root == nil {
		return
	}
	var node = root
	for node != nil {
		var tmpNode = TreeNode{0, nil, node.Right}
		if node.Left != nil {
			node.Right = node.Left
			node.Left = nil
			toRightEnd(node.Right, tmpNode.Right)
		}
		// 如何把最右侧的挂载到Left的后面
		node = node.Right
	}
}

func toRightEnd(node *TreeNode, originalRightNode *TreeNode) {
	if node == nil {
		return
	} else if node.Right == nil {
		node.Right = originalRightNode
		return
	} else {
		toRightEnd(node.Right, originalRightNode)
	}
}
