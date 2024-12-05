package problem_200

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	footMap := make([][]bool, len(grid))
	for i := range footMap {
		footMap[i] = make([]bool, len(grid[0]))
	}

	// 深度搜索，从左往右，从上到下
	for i := range grid {
		for j := range grid[i] {
			res += deepSearch(grid, footMap, i, j)
		}
	}
	for i := range footMap {
		fmt.Println(footMap[i])
	}
	return res
}

func deepSearch(grid [][]byte, footMap [][]bool, i int, j int) int {
	iMax := len(grid)
	jMax := len(grid[0])
	if !footMap[i][j] && grid[i][j] == '1' {
		// 展开深度遍历
		footMap[i][j] = true
		if i-1 >= 0 {
			deepSearch(grid, footMap, i-1, j)
		}
		if i+1 < iMax { // 向下
			deepSearch(grid, footMap, i+1, j)
		}
		if j-1 >= 0 {
			deepSearch(grid, footMap, i, j-1)
		}
		if j+1 < jMax {
			deepSearch(grid, footMap, i, j+1)
		}
		return 1
	} else {
		// 终止了 或者是没找到
		return 0
	}
}
