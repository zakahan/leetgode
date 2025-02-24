// -------------------------------------------------
// Package problem_448
// Author: hanzhi
// Date: 2025/1/31
// -------------------------------------------------

package problem_448

func findDisappearedNumbers1(nums []int) []int {
	myMap := make(map[int]int)
	for _, n := range nums {
		myMap[n] += 1
	}
	res := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if myMap[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
