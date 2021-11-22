package top_question

/// Find the Duplicate Number
// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only constant extra space.
// Input: nums = [1,3,4,2,2]
// Output: 2
// Input: nums = [3,1,3,4,2]
// Output: 3
// Input: nums = [1,1]
// Output: 1
// Input: nums = [1,1,2]
// Output: 1

// From 142: Floyd's cycle detection algorithm
func FindDuplicate(nums []int) int {
	first := nums[0]
	t := nums[first]
	h := nums[t]
	for t != h {
		t, h = nums[t], nums[nums[h]]
	}
	t = nums[0]
	for t != h {
		t, h = nums[t], nums[h]
	}
	return h
}

// Modifying the array
// func FindDuplicate(nums []int) int {
// 	next := 0
// 	for nums[0] != nums[nums[0]] {
// 		next = nums[nums[0]]
// 		nums[nums[0]] = nums[0]
// 		nums[0] = next
// 	}
// 	return nums[0]
// }
