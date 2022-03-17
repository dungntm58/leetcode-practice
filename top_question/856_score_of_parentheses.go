package top_question

/// Score of Parentheses
// Given a balanced parentheses string s, return the score of the string.
// The score of a balanced parentheses string is based on the following rule:
// - "()" has score 1.
// - AB has score A + B, where A and B are balanced parentheses strings.
// - (A) has score 2 * A, where A is a balanced parentheses string.
// Input: s = "()"
// Output: 1
// Input: s = "(())"
// Output: 2
// Input: s = "()()"
// Output: 2

func ScoreOfParentheses(s string) int {
	res, balance := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			balance++
		} else {
			balance--
			if s[i-1] == '(' {
				res += 1 << balance
			}
		}
	}
	return res
}
