package top_question

/// Rotate Array
// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

func RotateArray(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}
	for i := 0; i < (len(nums)-k)/2; i++ {
		nums[i+k], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i+k]
	}
}

// func RotateArray(nums []int, k int) {
// 	k %= len(nums)
// 	if k == 0 {
// 		return
// 	}
// 	extraNums := make([]int, k)
// 	for i := 0; i < k; i++ {
// 		extraNums[i] = nums[len(nums)-k+i]
// 	}
// 	for i := len(nums) - 1; i >= k; i-- {
// 		nums[i] = nums[i-k]
// 	}
// 	for i := 0; i < k; i++ {
// 		nums[i] = extraNums[i]
// 	}
// }
