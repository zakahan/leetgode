// -------------------------------------------------
// Package problem_234
// Author: hanzhi
// Date: 2025/1/2
// -------------------------------------------------

package problem_234

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fmt.Println("fast:", fast.Val)
	fmt.Println("slow:", slow.Val)
	// 此时slow必然是 中间节点或者中间前节点  也就是节点2 或者节点3
	half := reverseList(slow.Next)
	for head.Val == half.Val {
		head = head.Next
		half = half.Next
		if half == nil {
			return true
		}
	}

	//for half != nil {
	//	fmt.Println(strconv.Itoa(half.Val) + " ")
	//	half = half.Next
	//}
	return false
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}
