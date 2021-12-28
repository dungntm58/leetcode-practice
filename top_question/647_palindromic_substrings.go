package top_question

/// Palindromic Substrings
// Given a string s, return the number of palindromic substrings in it.
// A string is a palindrome when it reads the same backward as forward.
// A substring is a contiguous sequence of characters within the string.
// Input: s = "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
// Input: s = "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

func CountSubstrings(s string) int {
	count := 0
	for i := range s {
		len1, len2 := palindromeLength(s, i, i), palindromeLength(s, i, i+1)
		count += (len1+1)/2 + len2/2
	}
	return count
}
