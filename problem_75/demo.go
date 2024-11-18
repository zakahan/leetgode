package problem_75

func sortColorsTwo(nums []int) {
	recall := []int{0, 0, 0}
	//var i = 0
	for i := range nums {
		recall[nums[i]]++
	}

	var k = 0
	var i = 0

	for i < len(nums) {
		if recall[k] != 0 {
			nums[i] = k
			recall[k]--
			i++
		} else {
			k++
		}
	}
}
