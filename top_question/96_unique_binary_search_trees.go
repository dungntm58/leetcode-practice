package top_question

/// Given an integer n, return the number of structurally unique BST's (binary search trees)
// which has exactly n nodes of unique values from 1 to n.
// Input: n = 3
// Output: 5
// Input: n = 1
// Output: 1

// Formula: (2n)! / ((n + 1)! * n!)
func NumTrees(n int) int {
	r := 1
	for i := 1; i <= n; i++ {
		r *= (n + i)
		r /= i
	}
	r /= (n + 1)
	return r
}
