package problem_39

import (
	"fmt"
	"sort"
)

// > 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 `target` ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，
// >并以列表形式返回。你可以按 任意顺序 返回这些组合。candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

var resList [][]int

func CombinationSum(candidates []int, target int) [][]int {
	resList = [][]int{}
	// 就是顺着找把
	arr := []int{}
	// 堆candidates排序
	sort.Ints(candidates)
	backtracking(candidates, arr, target, 0, 0)

	return resList
}

func backtracking(candidates []int, nowArray []int, target int, totalNum int, start int) {
	// totalNum 是当前的总值

	for i := start; i < len(candidates); i++ {
		value := totalNum + candidates[i]
		if value > target {
			fmt.Println("nowArray更新，本次失败， 支路为:")
			for j := range nowArray {
				fmt.Print(nowArray[j])
			}
			fmt.Println(candidates[i])
			return // 直接退出了
		} else if value == target {
			newArray := make([]int, len(nowArray))
			copy(newArray, nowArray)
			newArray = append(newArray, candidates[i])
			resList = append(resList, newArray)
			return
		} else {
			nowArray = append(nowArray, candidates[i])
			totalNumNew := totalNum + candidates[i]
			backtracking(candidates, nowArray, target, totalNumNew, i)
			nowArray = nowArray[:len(nowArray)-1]

		}

	}
}
