package top_question

/// Subsets
// Given an integer array nums of unique elements, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// Input: nums = [0]
// Output: [[],[0]]

func Subsets(nums []int) [][]int {
	res := [][]int{}
	count := 1 << len(nums)
	for i := 0; i < count; i++ {
		tmp := []int{}
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) != 0 {
				tmp = append(tmp, nums[j])
			}
		}
		res = append(res, tmp)
	}
	return res
}
