package main

import (
	"fmt"
)

func main() {
	x := '0'
	y := '1'
	fmt.Println(x & y)
}

// i代表在前i个里面挑选  j表示总和
// 如今的刚好补上缺
// 前i-1个就满足
