package top_question

import "math"

/// 132 Pattern
// Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k]
// such that i < j < k and nums[i] < nums[k] < nums[j].
// Return true if there is a 132 pattern in nums, otherwise, return false.
// Input: nums = [1,2,3,4]
// Output: false
// Explanation: There is no 132 pattern in the sequence.
// Input: nums = [3,1,4,2]
// Output: true
// Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
// Input: nums = [-1,3,2,0]
// Output: true
// Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].

func Find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	stack := make([]int, 0)
	second := math.MinInt
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < second {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			second = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
