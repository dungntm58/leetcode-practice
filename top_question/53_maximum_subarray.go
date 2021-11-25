package top_question

/// Maximum Subarray
// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Input: nums = [1]
// Output: 1
// Input: nums = [5,4,-1,7,8]
// Output: 23

// Kadane's algorithm
func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	sum := 0
	ans := nums[0]
	for _, num := range nums {
		sum += num
		if sum < num {
			sum = num
		}
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}
