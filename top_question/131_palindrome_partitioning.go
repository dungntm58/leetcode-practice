package top_question

/// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.
// A palindrome string is a string that reads the same backward as forward.
// Input: s = "aab"
// Output: [["a","a","b"],["aa","b"]]
// Input: s = "a"
// Output: [["a"]]

func Partition(s string) [][]string {
	n := len(s)
	ans := [][]string{}
	dp := make([][]bool, n) // Indicate whether a substring of range row...col is a palindrome or not
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	stack := []string{}
	dfsParitionPalindrome(s, 0, &ans, &stack, dp)
	return ans
}

func dfsParitionPalindrome(s string, start int, ans *[][]string, stack *[]string, dp [][]bool) {
	if start >= len(s) {
		cloneStack := make([]string, len(*stack))
		copy(cloneStack, *stack)
		*ans = append(*ans, cloneStack)
		return
	}
	for i := start; i < len(s); i++ {
		if s[i] == s[start] && (i-start <= 2 || dp[start+1][i-1]) {
			dp[start][i] = true
			*stack = append(*stack, string(s[start:i+1]))
			dfsParitionPalindrome(s, i+1, ans, stack, dp)
			*stack = (*stack)[:len(*stack)-1]
		}
	}
}
