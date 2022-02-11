package top_question

/// Permutation in String
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

func CheckInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}
	s1Count, s2Count := make([]int, 26), make([]int, 26)
	for i := 0; i < s1Len; i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}
	count := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			count++
		}
	}
	for i := 0; i < s2Len-s1Len; i++ {
		if count == 26 {
			return true
		}
		s2Count[s2[i+s1Len]-'a']++
		if s2Count[s2[i+s1Len]-'a'] == s1Count[s2[i+s1Len]-'a'] {
			count++
		} else if s2Count[s2[i+s1Len]-'a'] == s1Count[s2[i+s1Len]-'a']+1 {
			count--
		}
		s2Count[s2[i]-'a']--
		if s2Count[s2[i]-'a'] == s1Count[s2[i]-'a'] {
			count++
		} else if s2Count[s2[i]-'a'] == s1Count[s2[i]-'a']-1 {
			count--
		}
	}
	return count == 26
}
