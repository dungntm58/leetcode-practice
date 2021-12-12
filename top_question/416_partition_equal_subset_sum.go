package top_question

/// Partition Equal Subset Sum
// Given a non-empty array nums containing only positive integers,
// find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.
// Input: nums = [1,5,11,5]
// Output: true
// Explanation: The array can be partitioned as [1, 5, 5] and [11].
// Input: nums = [1,2,3,5]
// Output: false
// Explanation: The array cannot be partitioned into equal sum subsets.

func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	ans := make([]bool, sum+1)
	ans[0] = true
	for _, num := range nums {
		for i := sum; i > 0; i-- {
			if i >= num {
				ans[i] = ans[i] || ans[i-num]
			}
		}
	}
	return ans[sum]
}
