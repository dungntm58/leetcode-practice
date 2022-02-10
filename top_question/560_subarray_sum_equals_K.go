package top_question

/// Subarray Sum Equals K
// Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.
// Input: nums = [1,1,1], k = 2
// Output: 2
// Input: nums = [1,2,3], k = 3
// Output: 2

func SubarraySum(nums []int, k int) int {
	count, sum := 0, 0
	numMap := make(map[int]int)
	numMap[0] = 1
	for _, num := range nums {
		sum += num
		if numMap[sum-k] > 0 {
			count += numMap[sum-k]
		}
		numMap[sum]++
	}
	return count
}

// All positive elements
// func SubarraySum(nums []int, k int) int {
// 	start, length, res, sum := 0, 1, 0, nums[0]
// 	for start+length <= len(nums) {
// 		if sum >= k {
// 			if sum == k {
// 				res++
// 			}
// 			sum -= nums[start]
// 			if start+length < len(nums) {
// 				sum += nums[start+length]
// 			} else if length > 1 {
// 				length--
// 			}
// 			start++
// 		} else {
// 			length++
// 			if start+length <= len(nums) {
// 				sum += nums[start+length-1]
// 			}
// 		}
// 	}
// 	return res
// }
