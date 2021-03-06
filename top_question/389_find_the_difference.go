package top_question

/// Find the Difference
// You are given two strings s and t.
// String t is generated by random shuffling string s and then add one more letter at a random position.
// Return the letter that was added to t.
// Input: s = "abcd", t = "abcde"
// Output: "e"
// Explanation: 'e' is the letter that was added.
// Input: s = "", t = "y"
// Output: "y"

func FindTheDifference(s string, t string) byte {
	letterArr := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		letterArr[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		letterArr[t[i]-'a']--
	}
	for i := 0; i < len(letterArr); i++ {
		if letterArr[i] > 0 {
			return byte(i + 'a')
		}
	}
	return 0
}

// func FindTheDifference(s string, t string) byte {
// 	var res byte
// 	for i := 0; i < len(s); i++ {
// 		res ^= s[i]
// 		res ^= t[i]
// 	}
// 	res ^= t[len(s)]
// 	return res
// }
