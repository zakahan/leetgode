package problem_17

import "fmt"

// 用队列
func LetterCombinations(digits string) []string {
	var queue []string = []string{}

	var signalData = [][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return signalData[int(digits[0]-'0')]
	}
	var recalls []int // 记录要使用哪些层,一次一次的
	for _, item := range digits {
		recalls = append(recalls, int(item-'0')) // recalls ={2,3}
	}
	queue = enQueue(queue, "")

	// 先将第一层加进去
	var head = ""
	for i := range recalls { // 只要recalls还不是空的
		var codes []string = signalData[recalls[i]]
		// 先出队列，找到元素然后组合
		n := len(queue)
		for n > 0 {
			queue, head = deQueue(queue)
			for _, s := range codes {
				queue = enQueue(queue, head+s)
			} // 入队列完成
			n--
		}
		fmt.Printf("对数字 %d 当前结果如下\n", recalls[i])
		fmt.Println(queue)
	}
	return queue

}

func enQueue(queue []string, elem string) []string {
	queue = append(queue, elem)
	return queue
}

func deQueue(queue []string) ([]string, string) {
	if len(queue) == 0 {
		return queue, ""
	}
	head := queue[0]
	queue = queue[1:]
	return queue, head
}
