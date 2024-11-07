package problem_22

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
type PareNode struct {
	NodeStr     string
	LeftRemain  int // 还剩几个左边的没用
	RightRemain int // 还剩几个右边的没用
}

func GenerateParenthesis(n int) []string {
	/*
	   	   ( -> ((
	   	       (( -> (()
	   	             (((
	       ( -> ()
	*/
	if n == 1 {
		return []string{"()"}
	}
	var queue []PareNode = []PareNode{}
	var queueCopy []PareNode = []PareNode{}
	queue = enQueue(queue, PareNode{"(", n - 1, n})
	var queueN = len(queue)
	for true {
		var flag = 0
		queueCopy = []PareNode{}
		// 只要queueN的值比0大（队列里还有东西
		// 首先出队列，然后处理一下，再入队列
		for i := 0; i < queueN; i++ {
			var node = PareNode{"", n, n}
			queue, node = deQueue(queue)
			queueCopy = enQueue(queueCopy, node) // 对队列做一次深拷贝

			if node.LeftRemain > 0 {
				var leftNode = PareNode{node.NodeStr + "(", node.LeftRemain - 1, node.RightRemain}
				queue = enQueue(queue, leftNode)
				flag = 1
			}
			if node.RightRemain > 0 && node.LeftRemain < node.RightRemain {
				// 只有剩余的右边的  比  剩余的左边的  多的时候能用

				var rightNode = PareNode{node.NodeStr + ")", node.LeftRemain, node.RightRemain - 1}
				queue = enQueue(queue, rightNode)
				flag = 1
			}

		}
		if flag == 0 {
			break
		}
		queueN = len(queue)
		// 一切结束
		// 判断离开条件 如果node的节点全为0了，自然就推出了
	}
	resList := make([]string, len(queueCopy))
	for i := range queueCopy {
		resList[i] = queueCopy[i].NodeStr
	}

	return resList
}

func enQueue(queue []PareNode, elem PareNode) []PareNode {
	queue = append(queue, elem)
	return queue
}

func deQueue(queue []PareNode) ([]PareNode, PareNode) {
	if len(queue) == 0 {
		return queue, PareNode{"", 0, 0}
	}
	head := queue[0]
	queue = queue[1:]
	return queue, head
}
