package three_sum

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	// 先对nums去重
	//重复最多有三个
	cNums := makeMap2Slice(nums)
	// 临时存储的哈希表
	resMap := make(map[string][]int)
	// 存储结果的表
	var resList [][]int
	// 写一个哈希表，存储每个值以及出现的次数。
	// 要多次copy了 不过这个无所谓了
	myMap := makeMap(cNums)
	for existItem, value := range cNums {
		myMap[value] -= 1 // 先去掉这个
		for i := range cNums {
			// value + cNums[i] + ? == 0 这就是条件，我要找出这个？是啥
			if i == existItem {
				// 这个变量得排除掉
			} else {
				myMap[cNums[i]] -= 1 // 哈希表里再去掉这个
				// 检查哈希表中是否存在对应的值
				if myMap[0-value-cNums[i]] > 0 {
					// 如果存在，那就可以了
					var tmp []int = []int{value, cNums[i], 0 - value - cNums[i]} // todo
					//resList = append(resList, tmp)
					resMap[code2Str(tmp)] = tmp
				}
				// 收尾工作
				myMap[cNums[i]] += 1
			}

		}
		// 收尾工作（对第一个数，补回来）
		myMap[value] += 1
	}
	for _, list := range resMap {
		resList = append(resList, list)
	}

	return resList
}

func makeMap(nums []int) map[int]int {
	myMap := make(map[int]int)
	for _, value := range nums {
		myMap[value] += 1
	}
	return myMap
}

// 去重，超过三个绝对就没用了。所以最多保留三个
func makeMap2Slice(nums []int) []int {
	myMap := make(map[int]int)
	for _, value := range nums {
		if myMap[value] < 3 {
			myMap[value] += 1
		}

	}
	var resNums []int
	for key, value := range myMap {
		for i := 0; i < value; i++ {
			resNums = append(resNums, key)
		}
	}
	return resNums

}

func code2Str(a []int) string {
	sort.Ints(a)
	result := ""
	for _, num := range a {
		result += fmt.Sprint("%d", num)
	}
	return result

}
