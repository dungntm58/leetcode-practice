package top_question

/// Detect Capital
// We define the usage of capitals in a word to be right when one of the following cases holds:
// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Given a string word, return true if the usage of capitals in it is right.
// Input: word = "USA"
// Output: true
// Input: word = "FlaG"
// Output: false

func DetectCapitalUse(word string) bool {
	uppercaseCount := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			uppercaseCount++
		}
	}
	if uppercaseCount == len(word) || uppercaseCount == 0 {
		return true
	}
	return uppercaseCount == 1 && word[0] >= 'A' && word[0] <= 'Z'
}
