package top_question

/// Delete and Earn
// You are given an integer array nums.
// You want to maximize the number of points you get by performing the following operation any number of times:
// - Pick any nums[i] and delete it to earn nums[i] points.
//   Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.
// Return the maximum number of points you can earn by applying the above operation some number of times.
// Input: nums = [3,4,2]
// Output: 6
// Explanation: You can perform the following operations:
// - Delete 4 to earn 4 points. Consequently, 3 is also deleted. nums = [2].
// - Delete 2 to earn 2 points. nums = [].
// You earn a total of 6 points.
// Input: nums = [2,2,3,3,3,4]
// Output: 9
// Explanation: You can perform the following operations:
// - Delete a 3 to earn 3 points. All 2's and 4's are also deleted. nums = [3,3].
// - Delete a 3 again to earn 3 points. nums = [3].
// - Delete a 3 once more to earn 3 points. nums = [].
// You earn a total of 9 points.

func DeleteAndEarn(nums []int) int {
	numMap, max := make(map[int]int), 0
	for _, num := range nums {
		numMap[num] += num
		if num > max {
			max = num
		}
	}
	dp := make([]int, max+1)
	dp[1] = numMap[1]
	for i := 2; i < max; i++ {
		prevPrev := numMap[i] + dp[i-2]
		if dp[i-1] > prevPrev {
			dp[i] = dp[i-1]
		} else {
			dp[i] = prevPrev
		}
	}
	return dp[max]
}
