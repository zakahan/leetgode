package problem_155

//
//type LinkNode struct {
//	Val  int
//	Next *LinkNode
//	Next *LinkNode
//}
//
//func addLinkNode(frontNode *LinkNode, itemNode *LinkNode) {
//	// frontNodes是插入点的前驱
//	rearNode := frontNode.Next
//	// 首先让itemNode指向两者
//	itemNode.Next = frontNode
//	itemNode.Next = rearNode
//	// 然后让他们的前后驱指向itemNode
//	rearNode.Next = itemNode
//	frontNode.Next = itemNode
//
//}
//
//func deleteNode(itemNode *LinkNode) *LinkNode {
//	frontNode := itemNode.Next
//	rearNode := itemNode.Next
//	// 解除关系
//	frontNode.Next = rearNode
//	rearNode.Next = frontNode
//	// item
//	itemNode.Next = nil
//	itemNode.Next = nil
//	return itemNode
//}
//
//type StackItem struct {
//	Node *LinkNode
//}
//
//type MinStack struct {
//	// 一个顺序栈、一个链表（从小到大）
//	Stack    []StackItem
//	LinkHead *LinkNode
//	LinkEnd  *LinkNode
//	stackTop int
//}
//
//func Constructor() MinStack {
//	stack := make([]StackItem, 30000)
//	headNode := LinkNode{Val: math.MinInt32, Next: nil, Next: nil}
//	endNode := LinkNode{Val: math.MaxInt32, Next: nil, Next: nil}
//	headNode.Next = &endNode
//	endNode.Next = &headNode
//	return MinStack{Stack: stack, LinkHead: &headNode, LinkEnd: &endNode, stackTop: -1}
//}
//
//func (this *MinStack) Push(val int) {
//	// 入栈
//	// 先声明元素
//	newNode := LinkNode{Val: val, Next: nil, Next: nil}
//	stackItem := StackItem{Node: &newNode}
//	this.stackTop++
//	this.Stack[this.stackTop] = stackItem
//	q := this.LinkHead
//	p := this.LinkHead.Next
//	for q != this.LinkEnd {
//		if q.Val <= newNode.Val && newNode.Val <= p.Val {
//			// 一旦在q和p数值之间 就插入
//			addLinkNode(q, &newNode)
//			return
//		}
//		q = q.Next
//		p = p.Next
//	}
//}
//
//func (this *MinStack) Pop() {
//	itemNode := this.Stack[this.stackTop].Node
//	this.stackTop--
//	itemNode = deleteNode(itemNode) // 删除节点
//}
//
//func (this *MinStack) Top() int {
//	return this.Stack[this.stackTop].Node.Val
//}
//
//func (this *MinStack) GetMin() int {
//	return this.LinkHead.Next.Val
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
