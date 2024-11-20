package problem_105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) > 0 {
		root := &TreeNode{preorder[0], nil, nil}
		var leftInorder []int
		var rightInorder []int
		for i := range inorder {
			if inorder[i] == preorder[0] {
				// 证明此时找到了
				leftInorder = inorder[:i]
				rightInorder = inorder[i+1:]
			}
		}
		// 左子树
		root.Left = buildTree(preorder[1:1+len(leftInorder)], leftInorder)
		root.Right = buildTree(preorder[1+len(leftInorder):], rightInorder)
		return root
	} else {
		return nil
	}
}
