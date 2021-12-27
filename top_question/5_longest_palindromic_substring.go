package top_question

/// Longest Palindromic Substring
// Given a string s, return the longest palindromic substring in s.
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Input: s = "cbbd"
// Output: "bb"

func LongestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1, len2 := palindromeLength(s, i, i), palindromeLength(s, i, i+1)
		len := len1
		if len < len2 {
			len = len2
		}
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return string(s[start : end+1])
}

func palindromeLength(s string, centerLeft, centerRight int) int {
	l, r := centerLeft, centerRight
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
