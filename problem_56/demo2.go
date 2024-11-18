package problem_56

import "sort"

func Merge(intervals [][]int) [][]int {
	var resList [][]int

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var j = -1
	for i := range intervals { // 新来的

		if len(resList) == 0 || resList[j][1] < intervals[i][0] {
			resList = append(resList, intervals[i])
			j++
		} else {
			if intervals[i][1] > resList[j][1] {
				resList[j][1] = intervals[i][1]
			}
		}

	}
	return resList
}
