package top_question

/// Product of Array Except Self
// Given an integer array nums,
// return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	leftProduct := make([]int, n)
	leftProduct[0] = nums[0]
	for i := 1; i < n; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i]
	}
	rightProduct := make([]int, n)
	rightProduct[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i]
	}
	r := make([]int, n)
	r[0] = rightProduct[1]
	r[n-1] = leftProduct[n-2]
	for i := 1; i < n-1; i++ {
		r[i] = leftProduct[i-1] * rightProduct[i+1]
	}
	return r
}
