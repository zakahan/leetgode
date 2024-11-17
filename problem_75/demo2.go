package problem_75

func sortColors(nums []int) {

	var p0 = 0
	var p1 = 0

	for i := range nums {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i] // 因为此时p0的位置应该是一个1，这个1得给还回去
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
