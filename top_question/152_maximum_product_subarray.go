package top_question

/// Maximum Product Subarray
// Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product,
// and return the product.
// It is guaranteed that the answer will fit in a 32-bit integer.
// A subarray is a contiguous subsequence of the array.
// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

func MaxProduct(nums []int) int {
	r, n := nums[0], len(nums)
	left, right := 1, 1
	for i, num := range nums {
		left *= num
		right *= nums[n-1-i]
		maxLR := left
		if maxLR < right {
			maxLR = right
		}
		if maxLR > r {
			r = maxLR
		}
		if left == 0 {
			left = 1
		}
		if right == 0 {
			right = 1
		}
	}
	return r
}
