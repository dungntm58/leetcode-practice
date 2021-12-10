package top_question

/// Domino and Tromino Tiling
// You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.
// Given an integer n, return the number of ways to tile an 2 x n board. Since the answer may be very large, return it modulo 109 + 7.
// In a tiling, every square must be covered by a tile.
// Two tilings are different if and only if there are two 4-directionally adjacent cells
// on the board such that exactly one of the tilings has both squares occupied by a tile.
// Input: n = 3
// Output: 5
// Explanation: The five different ways are show above.
// Input: n = 1
// Output: 1

// Formula: f(k) = 2 * f(k-1) + f(k-3)
func NumTilings(n int) int {
	mod := 1_000_000_007
	if n <= 2 {
		return n
	}
	curr, prev, ppre := 5, 2, 1
	for k := 4; k < n+1; k++ {
		tmp := prev
		prev = curr
		curr = (2*curr + ppre) % mod
		ppre = tmp
	}
	return curr
}
