package top_question

/// Remove Duplicate Letters
// Given a string s, remove duplicate letters so that every letter appears once and only once.
// You must make sure your result is the smallest in lexicographical order among all possible results.
// Input: s = "bcabc"
// Output: "abc"
// Input: s = "cbacdcbc"
// Output: "acdb"

func RemoveDuplicateLetters(s string) string {
	letterArr := [26]int{}
	visited := [26]bool{}
	for i := 0; i < len(s); i++ {
		letterArr[s[i]-'a']++
	}
	res := []byte{}
	for i := 0; i < len(s); i++ {
		letterArr[s[i]-'a']--
		if visited[s[i]-'a'] {
			continue
		}
		for len(res) > 0 && res[len(res)-1] > s[i] && letterArr[res[len(res)-1]-'a'] > 0 {
			visited[res[len(res)-1]-'a'] = false
			res = res[:len(res)-1]
		}
		res = append(res, s[i])
		visited[s[i]-'a'] = true
	}
	return string(res)
}
