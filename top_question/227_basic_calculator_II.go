package top_question

/// Basic Calculator II
// Given a string s which represents an expression, evaluate this expression and return its value.
// The integer division should truncate toward zero.
// You may assume that the given expression is always valid. All intermediate results will be in the range of [-231, 231 - 1].
// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().
// Input: s = "3+2*2"
// Output: 7
// Input: s = " 3/2 "
// Output: 1
// Input: s = " 3+5 / 2 "
// Output: 5

func Calculate_II(s string) int {
	r, lastNumber, currNumber := 0, 0, 0
	var operator byte = '+'
	for i := 0; i <= len(s); i++ {
		var char byte
		if i < len(s) {
			char = s[i]
		} else {
			char = '+'
		}
		if char == ' ' {
			continue
		}
		if char >= '0' && char <= '9' {
			currNumber = currNumber*10 + int(char) - '0'
			continue
		}
		switch operator {
		case '+':
			r += lastNumber
			lastNumber = currNumber
		case '-':
			r += lastNumber
			lastNumber = -currNumber
		case '*':
			lastNumber *= currNumber
		case '/':
			lastNumber /= currNumber
		}
		operator, currNumber = char, 0
	}
	r += lastNumber
	return r
}

// 3-2*2*1*0+3-2-1*3+3
