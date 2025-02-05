package problem_207

// 这题本质就是
// 就是看他是否要形成一个环路
// 我需要生成一个图，看看这个图有没有环路就行了
// 拓扑排序
func canFinishTP(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) <= 1 {
		return true
	}
	// 生成图
	myMap := genMap(prerequisites)
	var queue = []*Node{}

	//for _, node := range myMap {
	//	fmt.Print(node.Val, ":", node.InP, " | ")
	//}
	//fmt.Println()

	// 根据入度入队列
	for _, nodeArr := range myMap {
		if nodeArr.InP == 0 { // 一定是等于0，如果是-1代表被我用过了
			queue = append(queue, nodeArr)
		}
	}

	for len(queue) != 0 {
		nodeArr := queue[0]
		nodeArr.InP--
		queue = queue[1:]
		for _, nextNode := range nodeArr.Next {
			nextNode.InP--
			if nextNode.InP == 0 {
				queue = append(queue, nextNode)
			}
		}
	}

	// 打印展示
	//fmt.Println("////////////////////")
	//for _, node := range myMap {
	//	fmt.Print(node.Val, ":", node.InP, " | ")
	//}
	//fmt.Println()

	for _, nodeArr := range myMap {
		if nodeArr.InP != -1 { // 一定是等于0，如果是-1代表被我用过了
			return false
		}
	}
	return true

}

type Node struct {
	Next []*Node
	InP  int
}

func genMap(req [][]int) map[int]*Node {
	nodeHashMap := make(map[int]*Node)
	for i := range req {
		val0 := req[i][0]
		val1 := req[i][1]
		// 注意0号点和1号点，是  1->0的样式

		// 先看前驱是否存在，创建
		var node1 *Node
		var node0 *Node
		if nodeHashMap[val1] == nil {
			// 前驱不存在，创建
			node1 = &Node{Next: make([]*Node, 0)} // 不存在则创建
			nodeHashMap[val1] = node1
		} else {
			node1 = nodeHashMap[val1]
			nodeHashMap[val1] = node1
		}
		// 然后看后继，
		if nodeHashMap[val0] == nil {
			// 后继不存在，创建
			node0 = &Node{Next: make([]*Node, 0)}
			nodeHashMap[val0] = node0
		} else {
			node0 = nodeHashMap[val0]
			nodeHashMap[val0] = node0
		}
		// 然后将后继添加到前驱的Next的表中
		node1.Next = append(node1.Next, node0)
		node0.InP++

	}
	return nodeHashMap
}
