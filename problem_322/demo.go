// -------------------------------------------------
// Package problem_322
// Author: hanzhi
// Date: 2025/2/12
// -------------------------------------------------

package problem_322

import (
	"math"
)

// 这个策略宣告失败了，太慢了。 [1,2,5] 100这种就过不去了
var myMap map[int]int
var globalDeepth int

func coinChange0(coins []int, amount int) int {
	globalDeepth = math.MaxInt64 // 搞个贼大的值
	myMap = make(map[int]int)
	// 层序遍历+剪枝 但是怎么终止呢？
	backTrack(coins, amount, 0)
	if globalDeepth == math.MaxInt64 {
		return -1
	} else {
		return globalDeepth
	}
}

func backTrack(coins []int, target int, deepth int) {
	if target < 0 {
		// 代表这一层失败了
		return
	} else if target == 0 {
		if deepth < globalDeepth {
			globalDeepth = deepth // 修改最小深度
		}
		return
	} else { // 如果target >0说明还能继续
		if myMap[target] != 0 { // 剪枝
			return
		}
		// 逐个算
		for i := 0; i < len(coins); i++ {
			backTrack(coins, target-coins[i], deepth+1)
		}
	}
}
