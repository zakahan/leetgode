package problem_33

func Search(nums []int, target int) int {
	// nums已经旋转过了 就是形如 [k-1 ~ n, 0 ~ k]
	// 二分查找， 一定能分出 有序和含逆两部分， 如果mid 大于k，那么左侧就是有序，如果小于k那么左侧就是无序
	// 反之右侧就是无序
	// 加入一些特例分析
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		} else {
			return -1
		}
	}

	if nums[0] > nums[len(nums)-1] {
		return findAbnormal(nums, 0, len(nums)-1, target)
	} else {
		return findNormal(nums, 0, len(nums)-1, target)
	}

}

func findNormal(nums []int, start int, end int, target int) int {
	if start < 0 || end > len(nums)-1 {
		return -1
	}
	if start > end { // 证明找不到了
		return -1
	}
	// 这个是对有序的表做搜索，他一定是  x ~ y  逐渐递增的
	if target > nums[end] || target < nums[start] {
		// 证明失败了
		return -1
	} else {
		// 有成功可能性
		value := nums[(start+end)/2]
		if value == target {
			return (start + end) / 2
		} else if value > target {
			// 如果中间值比目标大，那么在前段
			return findNormal(nums, start, (start+end)/2-1, target)
		} else {
			return findNormal(nums, (start+end)/2+1, end, target)
		}
	}
}

func findAbnormal(nums []int, start int, end int, target int) int {
	if start < 0 || end > len(nums)-1 {
		return -1
	}
	if start > end {
		return -1 // 找不到了
	}
	// 4~7 0~2 target 6
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if target < nums[start] && target > nums[end] {
		return -1
	} else {
		if nums[mid] > nums[end] { // 左侧正常 右侧异常
			if target < nums[mid] && target >= nums[start] {
				// 数据在左侧
				return findNormal(nums, start, mid-1, target)
			} else {
				// 数据在异常侧（右侧）
				return findAbnormal(nums, mid+1, end, target)
			}

		} else { // 右侧正常，左侧异常
			if nums[mid] < target && target <= nums[end] {
				// 在右侧
				return findNormal(nums, mid+1, end, target)
			} else {
				// 在左侧
				return findAbnormal(nums, start, mid-1, target)
			}

		}
	}

}
