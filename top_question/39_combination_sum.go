package top_question

/// Combination Sum
// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen numbers sum to target.
// You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.
// It is guaranteed that the number of unique combinations
// that sum up to target is less than 150 combinations for the given input.
// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.
// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]
// Input: candidates = [2], target = 1
// Output: []

func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	temp := []int{}
	recursionCombinationSum(candidates, 0, target, &temp, &result)
	return result
}

func recursionCombinationSum(candidates []int, i int, target int, temp *[]int, result *[][]int) {
	if target < 0 || i >= len(candidates) {
		return
	}
	if target == 0 {
		cloned := make([]int, len(*temp))
		copy(cloned, *temp)
		*result = append(*result, cloned)
		return
	}
	*temp = append(*temp, candidates[i])
	recursionCombinationSum(candidates, i, target-candidates[i], temp, result)
	*temp = (*temp)[:len(*temp)-1]
	recursionCombinationSum(candidates, i+1, target, temp, result)
}
