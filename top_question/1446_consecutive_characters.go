package top_question

/// Consecutive Characters
// The power of the string is the maximum length of a non-empty substring that contains only one unique character.
// Given a string s, return the power of s.
// Input: s = "leetcode"
// Output: 2
// Explanation: The substring "ee" is of length 2 with the character 'e' only.
// Input: s = "abbcccddddeeeeedcba"
// Output: 5
// Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
// Input: s = "triplepillooooow"
// Output: 5
// Input: s = "hooraaaaaaaaaaay"
// Output: 11
// Input: s = "tourist"
// Output: 1

func MaxPower(s string) int {
	r := 1
	currCount := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			currCount++
			if currCount > r {
				r = currCount
			}
		} else {
			currCount = 1
		}
	}
	return r
}
