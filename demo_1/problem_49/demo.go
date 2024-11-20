package problem_49

import (
	"strconv"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	myMap := make(map[string][]string) // 先做一个哈希表

	size := len(strs)
	cols := 26
	matrix := make([][]int, size)
	// 生成一个矩阵 用来存key-value
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i, sk := range strs { // sk string -> rune ,i 表示第i个字符串
		for _, s := range sk { // s是一个run
			index := int(s - 'a')
			matrix[i][index] += 1

		}
	}

	for i, _ := range matrix {
		strKey := make([]string, cols)
		for j, values := range matrix[i] {
			strKey[j] = strconv.Itoa(values)
		}
		strKey2 := strings.Join(strKey, ",")
		_, exist := myMap[strKey2]
		if exist {
			myMap[strKey2] = append(myMap[strKey2], strs[i])
		} else {
			myMap[strKey2] = []string{strs[i]}
		}
	}

	// 创建一个切片来存储 map 的值
	values := make([][]string, 0, len(myMap))

	// 遍历 map，将每个值添加到切片中
	for _, value := range myMap {
		values = append(values, value)
	}
	return values
}
