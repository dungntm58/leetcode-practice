package top_question

/// Burst Balloons
// You are given n balloons, indexed from 0 to n - 1.
// Each balloon is painted with a number on it represented by an array nums.
// You are asked to burst all the balloons.
// If you burst the ith balloon, you will get nums[i - 1] * nums[i] * nums[i + 1] coins.
// If i - 1 or i + 1 goes out of bounds of the array, then treat it as if there is a balloon with a 1 painted on it.
// Return the maximum coins you can collect by bursting the balloons wisely.
// Input: nums = [3,1,5,8]
// Output: 167
// Explanation:
// nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
// coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
// Input: nums = [1,5]
// Output: 10

func MaxCoins(nums []int) int {
	n := len(nums) + 2 // include -1 and n
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	newNums := make([]int, n)
	newNums[0] = 1
	newNums[n-1] = 1
	for i, num := range nums {
		newNums[i+1] = num
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			for k := i + 1; k < j; k++ {
				r := dp[i][k] + dp[k][j] + newNums[i]*newNums[k]*newNums[j]
				if dp[i][j] < r {
					dp[i][j] = r
				}
			}
		}
	}
	return dp[0][n-1]
}
