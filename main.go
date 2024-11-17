package main

import (
	"fmt"
	"goleetcode/problem_76"
)

func main() {
	fmt.Println("Calling Tes	t1 function from 11121.go")
	var s = "a"
	var t = "b"
	problem_76.MinWindow(s, t)

}

// i代表在前i个里面挑选  j表示总和
// 如今的刚好补上缺
// 前i-1个就满足
