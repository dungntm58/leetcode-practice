package top_question

/// Smallest Integer Divisible by K
// Given a positive integer k, you need to find the length of the smallest positive integer n such that n is divisible by k,
// and n only contains the digit 1.
// Return the length of n. If there is no such n, return -1.
// Note: n may not fit in a 64-bit signed integer.
// Input: k = 1
// Output: 1
// Explanation: The smallest answer is n = 1, which has length 1.
// Input: k = 2
// Output: -1
// Explanation: There is no such positive integer n divisible by 2.
// Input: k = 3
// Output: 3
// Explanation: The smallest answer is n = 111, which has length 3.

func SmallestRepunitDivByK(k int) int {
	remainder := 0
	for i := 1; i <= k; i++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return i
		}
	}
	return -1
}
