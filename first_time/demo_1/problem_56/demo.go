package problem_56

import "sort"

func merge(intervals [][]int) [][]int {
	var resList [][]int

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := range intervals { // 新来的
		mergeFlag := 0
		for j := range resList { // 旧有的
			// 找这个是否在resList里面存在
			// res |      int <
			// 在里面 |----<--->-----|
			if resList[j][0] <= intervals[i][0] && intervals[i][1] <= resList[j][1] {
				resList[j][0] = resList[j][0]
				resList[j][1] = resList[j][1]
				mergeFlag = 1
				break
			} else if intervals[i][0] <= resList[j][0] && intervals[i][1] <= resList[j][1] && resList[j][0] <= intervals[i][1] {
				//  这种是  < --- | --- > |
				resList[j][0] = intervals[i][0]
				resList[j][1] = resList[j][1]
				mergeFlag = 1
				break
			} else if resList[j][0] <= intervals[i][0] && resList[j][1] <= intervals[i][1] && intervals[i][0] <= resList[j][1] {
				//  这种是  | --- < --- | -- >
				resList[j][0] = resList[j][0]
				resList[j][1] = intervals[i][1]
				mergeFlag = 1
				break
			} else if intervals[i][0] <= resList[j][0] && resList[j][1] <= intervals[i][1] {
				// < -- | -- | <
				resList[j][0] = intervals[i][0]
				resList[j][1] = intervals[i][1]
				mergeFlag = 1
				break
			}
		} // resList里都没有查找到的话
		if mergeFlag == 0 {
			resList = append(resList, intervals[i])
		}

	}
	return resList
}
