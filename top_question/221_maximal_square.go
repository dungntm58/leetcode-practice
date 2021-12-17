package top_question

/// Maximal Square
// Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// Output: 4
// Input: matrix = [["0","1"],["1","0"]]
// Output: 1
// Input: matrix = [["0"]]
// Output: 0

// 1: d(i,j) = min(d(i-1,j), d(i,j-1), d(i-1,j-1)) + 1
// 2: prevDp, d(j) = d(j), min(d(j-1), d(j), prevDp) + 1
func MaximalSquare(matrix [][]byte) int {
	row, col := len(matrix), len(matrix[0])
	maxEdgeLen, prevDp := 0, 0
	dp := make([]int, col+1)
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			temp := dp[j]
			if matrix[i-1][j-1] == '1' {
				t := prevDp
				if t > dp[j-1] {
					t = dp[j-1]
				}
				if t > dp[j] {
					t = dp[j]
				}
				dp[j] = t + 1
				if maxEdgeLen < dp[j] {
					maxEdgeLen = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prevDp = temp
		}
	}
	return maxEdgeLen * maxEdgeLen
}
