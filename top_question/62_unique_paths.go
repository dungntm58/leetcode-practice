package top_question

/// Unique Paths
// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?
// Input: m = 3, n = 7
// Output: 28
// Input: m = 3, n = 2
// Output: 3
// Explanation:
// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Down -> Down
// 2. Down -> Down -> Right
// 3. Down -> Right -> Down
// Input: m = 7, n = 3
// Output: 28
// Input: m = 3, n = 3
// Output: 6

func UniquePaths(m int, n int) int {
	ans := make([]int, n)
	ans[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			ans[j] += ans[j-1]
		}
	}
	return ans[n-1]
}

// Formula: (m+n-2)! / (n-1)! (m-1)!
