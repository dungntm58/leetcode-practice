package top_question

/// Smallest String With A Given Numeric Value
// The numeric value of a lowercase character is defined as its position (1-indexed) in the alphabet,
// so the numeric value of a is 1,
// the numeric value of b is 2, the numeric value of c is 3, and so on.
// The numeric value of a string consisting of lowercase characters is defined as the sum of its characters' numeric values.
// For example, the numeric value of the string "abe" is equal to 1 + 2 + 5 = 8.
// You are given two integers n and k.
// Return the lexicographically smallest string with length equal to n and numeric value equal to k.
// Note that a string x is lexicographically smaller than string y if x comes before y in dictionary order,
// that is, either x is a prefix of y, or if i is the first position such that x[i] != y[i],
// then x[i] comes before y[i] in alphabetic order.
// Input: n = 3, k = 27
// Output: "aay"
// Explanation: The numeric value of the string is 1 + 1 + 25 = 27,
// and it is the smallest string with such a value and length equal to 3.
// Input: n = 5, k = 73
// Output: "aaszz"

func GetSmallestString(n int, k int) string {
	if n > k {
		return ""
	}
	simplifiedBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		simplifiedBytes[i] = byte(1)
	}
	capacity := k - n
	maxCharacterIndex := n - 1
	for capacity > 0 && maxCharacterIndex >= 0 {
		if capacity > 25 {
			simplifiedBytes[maxCharacterIndex] = 26
			capacity -= 25
		} else {
			simplifiedBytes[maxCharacterIndex] = byte(capacity)
			capacity = 0
		}
		maxCharacterIndex--
	}
	for i := 0; i < n; i++ {
		simplifiedBytes[i] += 'a' - 1
	}
	return string(simplifiedBytes)
}
