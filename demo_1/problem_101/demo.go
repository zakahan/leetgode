package problem_101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// 用队列法解决 层次遍历
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
				symList = append(symList, 101)
			}
			if node != nil {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
			i++
		}
		queue = queue[l:]
		// 每次结束后判断是否对称
		var n = len(symList)
		//fmt.Println("本次实验：")
		for j := 0; j < n; j++ {
			//fmt.Print(symList[j])
			//fmt.Print(" ")
			if symList[j] == symList[n-1-j] {
				// pass
			} else {
				return false
			}
		}
	}
	return true
}
