package top_question

/// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
// Input: s = ""
// Output: 0

func LengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	r := 0
	charMap := map[rune]int{}
	currLen := 0
	leftI := -1
	for i, char := range chars {
		if charI, ok := charMap[char]; ok && leftI < charI {
			leftI = charI
		}
		currLen = i - leftI
		if currLen > r {
			r = currLen
		}
		charMap[char] = i
	}
	currLen = len(chars) - leftI - 1
	if currLen > r {
		return currLen
	}
	return r
}
