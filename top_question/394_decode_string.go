package top_question

import (
	"strconv"
)

/// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
// Note that k is guaranteed to be a positive integer.
// You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
// Furthermore, you may assume that the original data does not contain any digits
// and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"
// Input: s = "3[a2[c]]"
// Output: "accaccacc"
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"
// Input: s = "abc3[cd]xyz"
// Output: "abccdcdcdxyz"

func DecodeString(s string) string {
	return decodeSubString([]rune(s), 1)
}

func decodeSubString(runes []rune, times int) string {
	startBracketIndex, openingBracketCount, endBracketIndex := -1, 0, -1
bracketLoop:
	for i, c := range runes {
		switch c {
		case '[':
			if startBracketIndex < 0 {
				startBracketIndex = i
			}
			openingBracketCount++
		case ']':
			endBracketIndex = i
			if openingBracketCount == 1 {
				break bracketLoop
			}
			openingBracketCount--
		}
	}

	number := 1
	numberIndex := startBracketIndex - 1
	if startBracketIndex > 0 {
		for numberIndex >= 0 && runes[numberIndex] <= '9' {
			numberIndex--
		}
		numberIndex++
		number, _ = strconv.Atoi(string(runes[numberIndex:startBracketIndex]))
	}
	var r string
	if numberIndex < 0 {
		r = string(runes)
	} else {
		r = string(runes[0:numberIndex])
	}
	if startBracketIndex > 0 {
		r += decodeSubString(runes[startBracketIndex+1:endBracketIndex], number)
		if endBracketIndex < len(runes) {
			r += decodeSubString(runes[endBracketIndex+1:], 1)
		}
	}
	if times == 1 {
		return r
	}
	finalStr := r
	for i := 1; i < times; i++ {
		finalStr += r
	}
	return finalStr
}
