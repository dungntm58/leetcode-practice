package top_question

/// Find All Anagrams in a String
// Given two strings s and p, return an array of all the start indices of p's anagrams in s.
// You may return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
// Input: s = "abab", p = "ab"
// Output: [0,1,2]
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".

func FindAnagrams(s string, p string) []int {
	res := []int{}
	pArr := make([]int, 26)
	for i := 0; i < len(p); i++ {
		pArr[p[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if pArr[i] == 0 {
			pArr[i] = -1
		}
	}
LOOP:
	for i := 0; i < len(s); {
		count := len(p)
		cloneArr := make([]int, 26)
		copy(cloneArr, pArr)
		for j := i; j < i+len(p) && j < len(s); j++ {
			switch cloneArr[s[j]-'a'] {
			case -1:
				i = j + 1
				continue LOOP
			case 0:
				i++
				continue LOOP
			default:
				cloneArr[s[j]-'a']--
				count--
			}
			if count == 0 {
				res = append(res, i)
				break
			}
		}
		i++
	}
	return res
}
