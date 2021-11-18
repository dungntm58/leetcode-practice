package top_question

/// Find All Numbers Disappeared in an Array
// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
// Input: nums = [1,1]
// Output: [2]

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	ans := make([]bool, n)
	count := 0
	for _, num := range nums {
		if !ans[num-1] {
			count++
			ans[num-1] = true
		}
	}
	if n == count {
		return []int{}
	}
	r := make([]int, n-count)
	j := 0
	for i, el := range ans {
		if !el {
			r[j] = i + 1
			j++
		}
	}
	return r
}
