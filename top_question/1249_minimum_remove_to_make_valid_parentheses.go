package top_question

/// Minimum Remove to Make Valid Parentheses
// Given a string s of '(' , ')' and lowercase English characters.
// Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions )
// so that the resulting parentheses string is valid and return any valid string.
// Formally, a parentheses string is valid if and only if:
// - It is the empty string, contains only lowercase characters, or
// - It can be written as AB (A concatenated with B), where A and B are valid strings, or
// - It can be written as (A), where A is a valid string.
// Input: s = "lee(t(c)o)de)"
// Output: "lee(t(c)o)de"
// Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
// Input: s = "a)b(c)d"
// Output: "ab(c)d"
// Input: s = "))(("
// Output: ""
// Explanation: An empty string is also valid.

func MinRemoveToMakeValid(s string) string {
	sArr := []byte(s)
	stack := []int{}
	res := []byte{}
	for i, c := range sArr {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			if len(stack) == 0 {
				sArr[i] = '*'
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for len(stack) > 0 {
		sArr[stack[len(stack)-1]] = '*'
		stack = stack[:len(stack)-1]
	}
	for _, c := range sArr {
		if c != '*' {
			res = append(res, c)
		}
	}
	return string(res)
}
