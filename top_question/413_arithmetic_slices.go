package top_question

/// Arithmetic Slices
// An integer array is called arithmetic if it consists of at least three elements
// and if the difference between any two consecutive elements is the same.
// For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.
// Given an integer array nums, return the number of arithmetic subarrays of nums.
// A subarray is a contiguous subsequence of the array.
// Input: nums = [1,2,3,4]
// Output: 3
// Explanation: We have 3 arithmetic slices in nums: [1, 2, 3], [2, 3, 4] and [1,2,3,4] itself.
// Input: nums = [1]
// Output: 0

func NumberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	res, start, end, diff := 0, 0, 0, nums[1]-nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[i-1]+diff {
			end = i
		} else {
			if end-start > 1 {
				res += (end - start) * (end - start - 1) / 2
			}
			start, end, diff = i-1, i, nums[i]-nums[i-1]
		}
	}
	if end-start > 1 {
		res += (end - start) * (end - start - 1) / 2
	}
	return res
}
