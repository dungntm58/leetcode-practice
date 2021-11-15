package top_question

import (
	"sort"
)

/// Largest Divisible Subset
// Given a set of distinct positive integers nums,
// return the largest subset answer such that every pair (answer[i], answer[j]) of elements in this subset satisfies:
// - answer[i] % answer[j] == 0, or
// - answer[j] % answer[i] == 0
// If there are multiple solutions, return any of them.
// Input: nums = [1,2,3]
// Output: [1,2]
// Explanation: [1,3] is also accepted.
// Input: nums = [1,2,4,8]
// Output: [1,2,4,8]

func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, n)
	ans[0] = []int{nums[0]}
	maxCount := 1
	maxCountIndex := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			var r []int
			if nums[i]%ans[j][len(ans[j])-1] == 0 {
				lenr := len(ans[j]) + 1
				r = make([]int, lenr)
				copy(r, ans[j])
				r[lenr-1] = nums[i]
			} else {
				r = []int{nums[i]}
			}
			if len(r) > maxCount {
				maxCount = len(r)
				maxCountIndex = i
			}
			if len(r) > len(ans[i]) {
				ans[i] = r
			}
		}
	}
	return ans[maxCountIndex]
}
