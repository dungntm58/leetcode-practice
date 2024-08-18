package top_question

// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
// Given an integer n, return the nth ugly number.

// Example 1:

// Input: n = 10
// Output: 12
// Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.
// Example 2:

// Input: n = 1
// Output: 1
// Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.

// Constraints:
// - 1 <= n <= 1690

func NthUglyNumberII(n int) int {
	var i2, i3, i5 int
	uglyArr := make([]int, n)
	uglyArr[0] = 1
	for i := 1; i < n; i++ {
		nextI2, nextI3, nextI5 := uglyArr[i2]*2, uglyArr[i3]*3, uglyArr[i5]*5
		uglyArr[i] = min(nextI2, nextI3, nextI5)
		if uglyArr[i] == nextI2 {
			i2++
		}
		if uglyArr[i] == nextI3 {
			i3++
		}
		if uglyArr[i] == nextI5 {
			i5++
		}
	}
	return uglyArr[n-1]
}
