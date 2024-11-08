package main

import (
	"fmt"
	"goleetcode/problem_23"
)

func main() {
	fmt.Println("Calling Test1 function from 11121.go")
	var list []*problem_23.ListNode = []*problem_23.ListNode{}
	var list11 = problem_23.ListNode{1, nil}
	var list21 = problem_23.ListNode{10, nil}
	list = append(list, &list11)
	list = append(list, &list21)
	//m := problem_23.MergeTwoList(&list11, &list21)
	//fmt.Println(m.Val)
}
