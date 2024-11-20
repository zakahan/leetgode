package main

import (
	"fmt"
	"goleetcode/demo_1/problem_78"
)

func main() {
	fmt.Println("Calling Tes	t1 function from 11121.go")
	var nums []int = []int{1, 2, 3}
	x := problem_78.Subsets(nums)
	for i := range x {
		fmt.Println(x[i])
	}
}

// i代表在前i个里面挑选  j表示总和
// 如今的刚好补上缺
// 前i-1个就满足
