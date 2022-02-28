package top_question

import "strconv"

/// Summary Ranges
// You are given a sorted unique integer array nums.
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
// That is, each element of nums is covered by exactly one of the ranges,
// and there is no integer x such that x is in one of the ranges but not in nums.
// Each range [a,b] in the list should be output as:
// - "a->b" if a != b
// - "a" if a == b
// Input: nums = [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
// Explanation: The ranges are:
// [0,2] --> "0->2"
// [4,5] --> "4->5"
// [7,7] --> "7"
// Input: nums = [0,2,3,4,6,8,9]
// Output: ["0","2->4","6","8->9"]
// Explanation: The ranges are:
// [0,0] --> "0"
// [2,4] --> "2->4"
// [6,6] --> "6"
// [8,9] --> "8->9"

func SummaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); {
		j := i + 1
		for j < len(nums) && nums[j] == nums[j-1]+1 {
			j++
		}
		if i == j-1 {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
		}
		i = j
	}
	return res
}
