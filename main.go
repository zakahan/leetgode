package main

import "fmt"

func main() {
	for i := 0; i < 10001; i++ {
		x := i * i
		if x <= 10001 {
			fmt.Printf("%v, ", x)
		}
	}
}

// i代表在前i个里面挑选  j表示总和
// 如今的刚好补上缺
// 前i-1个就满足
