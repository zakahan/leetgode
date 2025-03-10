// -------------------------------------------------
// Package classical_alg
// Author: hanzhi
// Date: 2025/3/10
// -------------------------------------------------

package classical_alg

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	arr := []int{
		8, 5, 8, 1, 3, 0, 3, 9, 2, 1, 7, 3, 10, 11, 19, 101, 4, 5,
	}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, height int) {
	if low >= height {
		return
	}
	pIndex := paritioin(arr, low, height)
	quickSort(arr, low, pIndex-1)
	quickSort(arr, pIndex+1, height)

}

func paritioin(arr []int, low, height int) int {
	pValue := arr[height]
	// 开始遍历
	i := low - 1
	for j := low; j < height; j++ {
		if arr[j] <= pValue { // 一定是当前值小于分列的，换到前面（i可以理解为前面）
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 补全
	arr[i+1], arr[height] = arr[height], arr[i+1]
	return i + 1
}
