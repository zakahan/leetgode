// -------------------------------------------------
// Package problem_189
// Author: hanzhi
// Date: 2025/2/26
// -------------------------------------------------

package problem_189

func rotate(nums []int, k int) {
	k = k % len(nums)
	// 两次反转
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	// 局部两次反转
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}

	for i := k; i < (k+n)/2; i++ {
		nums[i], nums[n-1-(i-k)] = nums[n-1-(i-k)], nums[i]
	}
}
