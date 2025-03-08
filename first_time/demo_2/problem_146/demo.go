package problem_146

import "fmt"

// nb! 五十个了
/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。  (这个感觉map最合适）
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/
type LinkStackNode struct {
	Key   int // 多存一个key吧，没办法，不然清空不了
	Val   int
	Next  *LinkStackNode
	Front *LinkStackNode
}

type LRUCache struct {
	myMap       map[int]*LinkStackNode // 用来辅助get操作的
	Capacity    int
	StackNumber int
	StackTop    *LinkStackNode
	StackDown   *LinkStackNode
}

func Constructor(capacity int) LRUCache {
	stackHead := LinkStackNode{Key: -111, Val: 159, Next: nil, Front: nil}
	stackEnd := LinkStackNode{Key: -111, Val: -159, Next: nil, Front: nil}
	stackHead.Next = &stackEnd
	stackEnd.Front = &stackHead
	return LRUCache{myMap: make(map[int]*LinkStackNode), Capacity: capacity, StackNumber: 0, StackTop: &stackHead, StackDown: &stackEnd}
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("-----------------------------------")
	printf(this.StackTop, "before get")
	var value = -1
	if this.myMap[key] != nil {
		thisNode := this.myMap[key]
		frontNode := this.myMap[key].Front
		nextNode := this.myMap[key].Next

		// 将thisNode放到栈顶
		// 1. 将原有的收尾链接
		frontNode.Next = nextNode
		nextNode.Front = frontNode
		// 2. 将自己放到栈顶
		preTopNode := this.StackTop.Next
		preTopNode.Front = thisNode
		this.StackTop.Next = thisNode
		thisNode.Front = this.StackTop
		thisNode.Next = preTopNode
		value = thisNode.Val

	}
	printf(this.StackTop, "after get")
	fmt.Println("get key-value: ", key, value)
	fmt.Println("-----------------------------------")
	return value
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println("-----------------------------------")
	printf(this.StackTop, "before Put")
	// 添加内容
	if this.myMap[key] == nil { // 不存在，要判断容量够不够
		// 如果容量还够，那么无所谓了，直接加到栈顶就行了
		if this.StackNumber+1 <= this.Capacity {
		} else {
			// 容量不够了，直接出栈一个就行了。
			//fmt.Println("出栈！！")
			// 出栈一个元素
			tmp := this.StackDown.Front // 待出栈元素
			preFrontNode := tmp.Front
			preFrontNode.Next = this.StackDown
			this.StackDown.Front = preFrontNode

			// 将map对应的key清空
			tmpKey := tmp.Key
			this.myMap[tmpKey] = nil
		}
		preTopNode := this.StackTop.Next
		thisNode := LinkStackNode{Key: key, Val: value, Next: this.StackTop.Next, Front: this.StackTop}
		preTopNode.Front = &thisNode
		this.StackTop.Next = &thisNode

		this.myMap[key] = &thisNode
		this.StackNumber++
	} else {
		//this.StackNumber++
		//fmt.Println("Put", this.StackNumber, this.Capacity)
		// 如果存在，那么就不需要判断容量，只需要创建个新的把这放到栈顶
		thisNode := this.myMap[key]
		thisNode.Val = value
		// 找出他对应的前驱和后继，连接起来
		frontNode := this.myMap[key].Front
		nextNode := this.myMap[key].Next
		frontNode.Next = nextNode
		nextNode.Front = frontNode

		// 将this node放到栈顶
		thisNode.Next = this.StackTop.Next
		thisNode.Front = this.StackTop
		this.StackTop.Next = thisNode
		this.myMap[key] = thisNode
	}
	printf(this.StackTop, "after Put")
	fmt.Println("-----------------------------------")
}

func printf(head *LinkStackNode, s string) {
	fmt.Println(s)
	x := head
	for x != nil {
		fmt.Print(x.Key, " ")
		x = x.Next
	}
	fmt.Println()

}
