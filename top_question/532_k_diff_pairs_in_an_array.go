package top_question

/// K-diff Pairs in an Array
// Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.
// A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:
// - 0 <= i < j < nums.length
// - |nums[i] - nums[j]| == k
// Notice that |val| denotes the absolute value of val.
// Input: nums = [3,1,4,1,5], k = 2
// Output: 2
// Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
// Although we have two 1s in the input, we should only return the number of unique pairs.
// Input: nums = [1,2,3,4,5], k = 1
// Output: 4
// Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
// Input: nums = [1,3,1,5,4], k = 0
// Output: 1
// Explanation: There is one 0-diff pair in the array, (1, 1).

func FindPairs(nums []int, k int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	res := 0
	for num, count := range numMap {
		if k > 0 && numMap[num+k] > 0 {
			res++
		}
		if k == 0 && count > 1 {
			res++
		}
	}
	return res
}
