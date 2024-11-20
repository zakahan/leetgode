package problem_102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 用队列法解决 层次遍历
	resList := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		var i = 0
		var l = len(queue)
		symList := make([]int, 0)
		for i < l {
			node := queue[i]
			if node != nil {
				symList = append(symList, node.Val)
			} else {
				//symList = append(symList, 101)
			}
			if node != nil {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
			i++
		}
		queue = queue[l:]
		// 每次结束后判断是否对称
		if len(symList) > 0 {
			resList = append(resList, symList)
		}
	}
	return resList
}
