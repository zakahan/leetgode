package main

import (
	"fmt"
	"goleetcode/problem_416"
)

func main() {
	fmt.Println("Calling Tes	t1 function from 11121.go")
	x := []int{1, 5, 11, 5}
	m := problem_416.CanPartition(x)
	print(m) // target 11
}

// i代表在前i个里面挑选  j表示总和
// 如今的刚好补上缺
// 前i-1个就满足
