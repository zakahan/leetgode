package problem_4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var n1 = len(nums1)
	var n2 = len(nums2)
	var left = (n1 + n2 + 1) / 2
	var right = (n1 + n2 + 2) / 2
	res := float64(getKth(nums1, nums2, 0, n1-1, 0, n2-1, left) + getKth(nums1, nums2, 0, n1-1, 0, n2-1, right))
	return res / 2
}

func getKth(nums1 []int, nums2 []int, n1I int, n1J int, n2I int, n2J int, k int) int {
	// 当前片段的长度
	var len1 = n1J - n1I + 1
	var len2 = n2J - n2I + 1
	// 统一长度 保证每次都是len1更短
	if len1 > len2 {
		return getKth(nums2, nums1, n2I, n2J, n1I, n1J, k)
	}
	if len1 == 0 {
		return nums2[n2I+k-1] // 从当前位置选k就好了
	}

	if k == 1 {
		return minInt(nums1[n1I], nums2[n2I]) // 如果是1 就这俩选项了，选小的呗
	}

	// 防止越界，选 k/2 或更短的长度
	part1 := minInt(k/2, len1)
	part2 := minInt(k/2, len2)
	locationK1 := n1I + part1 - 1
	locationK2 := n2I + part2 - 1

	if nums1[locationK1] > nums2[locationK2] {
		// 序列2更小
		return getKth(nums1, nums2, n1I, n1J, locationK2+1, n2J, k-part2)
	} else {
		// 是locationK1 + 1 因为当前值也不要了 处理的是“从locationK1之后开始的序列”
		return getKth(nums1, nums2, locationK1+1, n1J, n2I, n2J, k-part1)
	}
}

func minInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
