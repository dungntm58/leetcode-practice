package top_question

/// Kth Smallest Number in Multiplication Table
// Nearly everyone has used the Multiplication Table.
// The multiplication table of size m x n is an integer matrix mat where mat[i][j] == i * j (1-indexed).
// Given three integers m, n, and k, return the kth smallest element in the m x n multiplication table.
// Input: m = 3, n = 3, k = 5
// Output: 3
// Explanation: The 5th smallest number is 3.
// Input: m = 2, n = 3, k = 6
// Output: 6
// Explanation: The 6th smallest number is 6.

func FindKthNumber(m int, n int, k int) int {
	low, high := 1, m*n
	for low < high {
		mid := (low + high) / 2
		count := 0
		for i := 1; i <= m; i++ {
			if mid/i < n {
				count += mid / i
			} else {
				count += n
			}
		}
		if count >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
