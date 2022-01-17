package top_question

/// Word Pattern
// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
// Input: pattern = "abba", s = "dog cat cat fish"
// Output: false
// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false

func WordPattern(pattern string, s string) bool {
	wordMap := map[rune]string{}
	wordSet := map[string]struct{}{}
	currentLetterIndex := 0
	for _, c := range pattern {
		prev := currentLetterIndex
		if prev >= len(s) {
			return false
		}
		for currentLetterIndex < len(s) && s[currentLetterIndex] != ' ' {
			currentLetterIndex++
		}
		word := string(s[prev:currentLetterIndex])
		currentLetterIndex++
		if w, ok := wordMap[c]; ok {
			if w != word {
				return false
			}
		} else {
			if _, ok := wordSet[word]; ok {
				return false
			}
			wordMap[c] = word
			wordSet[word] = struct{}{}
		}
	}
	return currentLetterIndex >= len(s)
}
