package top_question

/// Longest Common Subsequence
// Given two strings text1 and text2, return the length of their longest common subsequence.
// If there is no common subsequence, return 0.
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted
// without changing the relative order of the remaining characters.
// - For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings.
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.

func LongestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	ans := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		ans[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				ans[i][j] = ans[i-1][j-1] + 1
			} else {
				ans[i][j] = ans[i-1][j]
				if ans[i][j] < ans[i][j-1] {
					ans[i][j] = ans[i][j-1]
				}
			}
		}
	}
	return ans[m][n]
}
