package main

import (
	"fmt"
	"goleetcode/problem_34"
)

func main() {
	fmt.Println("Calling Test1 function from 11121.go")
	nums := []int{5, 7, 7, 8, 8, 10}
	x := problem_34.SearchRange(nums, 858)
	fmt.Println(x)
}
