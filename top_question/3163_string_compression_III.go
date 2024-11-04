package top_question

import (
	"bytes"
	"strconv"
)

// Given a string word, compress it using the following algorithm:
// Begin with an empty string comp. While word is not empty, use the following operation:
// - Remove a maximum length prefix of word made of a single character c repeating at most 9 times.
// - Append the length of the prefix followed by c to comp.
// Return the string comp.

// Example 1:
// Input: word = "abcde"
// Output: "1a1b1c1d1e"
// Explanation:
// Initially, comp = "". Apply the operation 5 times, choosing "a", "b", "c", "d", and "e" as the prefix in each operation.
// For each prefix, append "1" followed by the character to comp.

// Example 2:
// Input: word = "aaaaaaaaaaaaaabb"
// Output: "9a5a2b"
// Explanation:
// Initially, comp = "". Apply the operation 3 times, choosing "aaaaaaaaa", "aaaaa", and "bb" as the prefix in each operation.
// For prefix "aaaaaaaaa", append "9" followed by "a" to comp.
// For prefix "aaaaa", append "5" followed by "a" to comp.
// For prefix "bb", append "2" followed by "b" to comp.

// Constraints:
// - 1 <= word.length <= 2 * 105
// - word consists only of lowercase English letters.

func CompressedString(word string) string {
	var buffer bytes.Buffer
	if len(word) == 1 {
		buffer.WriteString("1")
		buffer.WriteString(word)
		return buffer.String()
	}
	i, count := 0, 1
	char := word[0]
	for i < len(word) {
		if word[i] != char {
			buffer.WriteString(strconv.Itoa(count))
			buffer.WriteByte(char)
			char = word[i]
			count = 1
		} else if count >= 9 {
			buffer.WriteString(strconv.Itoa(count))
			buffer.WriteByte(char)
			count = 1
		}
		count++
		i++
	}
	return buffer.String()
}
