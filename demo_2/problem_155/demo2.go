package problem_155

// 官方解法，双栈
// 一个正常顺序栈，另一个存当前最小值
type MinStack struct {
	StackNormal []int
	StackMin    []int
	StackTop    int
}

func Constructor() MinStack {
	sn := make([]int, 30000)
	sm := make([]int, 30000)
	top := -1
	return MinStack{StackNormal: sn, StackMin: sm, StackTop: top}
}

func (this *MinStack) Push(val int) {
	// 入栈
	if this.StackTop == -1 {
		this.StackTop++
		this.StackNormal[this.StackTop] = val
		this.StackMin[this.StackTop] = val
	} else {
		lastMinVal := this.StackMin[this.StackTop]
		this.StackTop++
		this.StackNormal[this.StackTop] = val
		if val < lastMinVal {
			this.StackMin[this.StackTop] = val
		} else {
			this.StackMin[this.StackTop] = lastMinVal
		}
	}
}

func (this *MinStack) Pop() {
	this.StackTop--
}

func (this *MinStack) Top() int {
	return this.StackNormal[this.StackTop]
}

func (this *MinStack) GetMin() int {
	return this.StackMin[this.StackTop]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
