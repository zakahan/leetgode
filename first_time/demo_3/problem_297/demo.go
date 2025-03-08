// -------------------------------------------------
// Package problem_297
// Author: hanzhi
// Date: 2025/2/3
// -------------------------------------------------

package problem_297

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const nullNumber = 1001

func IntSliceToString(intSlice []int) string {
	// 将[]int中的每个元素转换为字符串，并存储在一个新的字符串切片中
	strSlice := make([]string, len(intSlice))
	for i, num := range intSlice {
		strSlice[i] = strconv.Itoa(num)
	}
	// 使用strings.Join将字符串切片中的所有元素连接成一个单独的字符串，以逗号作为分隔符
	return strings.Join(strSlice, ",")
}

func StringToIntSlice(input string) []int {
	// 使用strings.Split按逗号分割输入字符串，得到字符串切片
	strSlice := strings.Split(input, ",")
	// 创建一个等长的整型切片用于存储结果
	intSlice := make([]int, len(strSlice))

	// 遍历字符串切片，将每个字符串元素转换为整数
	for i, str := range strSlice {
		val, err := strconv.Atoi(str)
		if err != nil {
			// 如果转换过程中出现错误，返回错误
			return nil
		}
		intSlice[i] = val
	}
	return intSlice
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

var resList []int

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 用深度优先遍历
	resList = make([]int, 0)
	// 逐个插入
	dfs(root)
	x := IntSliceToString(resList)
	return x
}

// 采用先序遍历的方式来做吧
func dfs(root *TreeNode) {
	if root != nil {
		resList = append(resList, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	} else {
		// 不然的话就添加1001
		resList = append(resList, nullNumber)
	}

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := StringToIntSlice(data)
	// 然后反序列化
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == nullNumber {
			sp = sp[1:]
			return nil
		}
		val := sp[0]
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
	// 来自官方题解的反序列方式，挺有巧思的设计方式

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
