package problem_34

// 很像39题 组合总数哪个
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	i := biSearchLeft(nums, 0, len(nums)-1, float64(target)-0.5)
	j := biSearchRight(nums, 0, len(nums)-1, float64(target)+0.5)
	return []int{i, j}
}

// 先做个二分查找
func biSearchLeft(arr []int, start int, end int, target float64) int {
	if int(target+0.5) < arr[start] || int(target+0.5) > arr[end] {
		return -1
	}
	if start > end {
		return -1
	} else {
		mid := (start + end) / 2
		var value float64 = float64(arr[mid])
		var valueLeft float64 = 0.0
		if mid == 0 {
			valueLeft = valueLeft - 1
		} else {
			valueLeft = float64(arr[mid-1])
		}
		// 比较
		if valueLeft < target && target < value { // 	不可能相等的
			if value == target+0.5 {
				return mid
			} else {
				return -1
			}
		} else if value < target { // 右侧
			return biSearchLeft(arr, mid+1, end, target)

		} else {
			return biSearchLeft(arr, start, mid-1, target)
		}
	}
}

func biSearchRight(arr []int, start int, end int, target float64) int {
	if int(target-0.5) < arr[start] || int(target-0.5) > arr[end] {
		return -1
	}
	if start > end {
		return -1
	} else {
		mid := (start + end) / 2
		var value float64 = float64(arr[mid])
		var valueRight float64 = 0.0
		if mid == len(arr)-1 {
			valueRight = value + 1
		} else {
			valueRight = float64(arr[mid+1])
		}
		// 比较
		if value < target && target < valueRight { // 	不可能相等的
			if value == target-0.5 {
				return mid
			} else {
				return -1
			}
		} else if valueRight < target { // 右侧
			return biSearchRight(arr, mid+1, end, target)

		} else {
			return biSearchRight(arr, start, mid-1, target)
		}
	}
}
