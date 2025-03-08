package problem_98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	resList := midOrder(root, make([]int, 0))
	for i := 0; i < len(resList)-1; i++ {
		if resList[i] >= resList[i+1] { // 不允许右等于
			return false
		}
	}
	return true
}

func midOrder(node *TreeNode, resList []int) []int {
	if node != nil {
		resList = midOrder(node.Left, resList)
		resList = append(resList, node.Val)
		resList = midOrder(node.Right, resList)
	} else {
		return resList
	}
	return resList
}
