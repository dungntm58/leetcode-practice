package top_question

/// Longest Valid Parentheses
// Given a string containing just the characters '(' and ')',
// find the length of the longest valid (well-formed) parentheses substring.
// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".
// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".
// Input: s = ""
// Output: 0

func LongestValidParentheses(s string) int {
	res, stack := 0, []int{}
	stack = append(stack, -1)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else if res < i-stack[len(stack)-1] {
				res = i - stack[len(stack)-1]
			}
		}
	}
	return res
}
