package problem_2

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 原地操作
	// 我献给l1和l2找俩爹节点（实体的，我还是觉得实体的方便
	p := ListNode{0, l1}
	z := &p // z始终指向处理节点的前驱
	// 然后开始遍历  用两个指针
	x := l1
	y := l2
	var carry int = 0
	var number int = 0
	for x != nil && y != nil {
		number = (x.Val + y.Val + carry) % 10
		carry = (x.Val + y.Val + carry) / 10
		x.Val = number
		x = x.Next
		y = y.Next
		z = z.Next
	}

	for x != nil && y == nil {
		// 还可以继续
		number = (x.Val + carry) % 10
		carry = (x.Val + carry) / 10
		x.Val = number
		x = x.Next
		z = z.Next

	}
	if y != nil && x == nil {
		z.Next = y // 变更z的后继
		for y != nil {
			number = (y.Val + carry) % 10
			carry = (y.Val + carry) / 10
			y.Val = number
			y = y.Next
			z = z.Next
		}
	}

	if carry > 0 {
		newNode := ListNode{carry, nil}
		z.Next = &newNode
	}
	return p.Next

}
