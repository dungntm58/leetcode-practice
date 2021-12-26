package top_question

/// Basic Calculator
// Given a string s representing a valid expression,
// implement a basic calculator to evaluate it,
// and return the result of the evaluation.
// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().
// Input: s = "1 + 1"
// Output: 2
// Input: s = " 2-1 + 2 "
// Output: 3
// Input: s = "(1+(4+5+2)-3)+(6+8)"
// Output: 23

func Calculate(s string) int {
	runes := []rune(s)
	simpifyExp(runes, 0, len(s)-1, false)
	r, curr := 0, 0
	var operator rune = '+'
	for i := 0; i <= len(runes); i++ {
		var char rune
		if i < len(runes) {
			char = runes[i]
		} else {
			char = '+'
		}
		if char == ' ' {
			continue
		}
		if char >= '0' && char <= '9' {
			curr = curr*10 + int(char) - '0'
			continue
		}
		if operator == '+' {
			r += curr
		} else {
			r -= curr
		}
		operator, curr = char, 0
	}
	return r
}

func simpifyExp(runes []rune, start, end int, reverse bool) {
	if start < 0 || end >= len(runes) || start >= end {
		return
	}
	openingBracketIndex, closingBracketIndex, openingBracketCount := -1, -1, 0
	for i := start; i <= end; i++ {
		if runes[i] == '(' {
			if openingBracketIndex < 0 {
				openingBracketIndex = i
			}
			openingBracketCount++
		} else if runes[i] == ')' {
			openingBracketCount--
			if openingBracketCount == 0 {
				closingBracketIndex = i
				break
			}
		}
	}
	reverseNextLevel := false
	if openingBracketIndex > 0 {
	checkMinusSign:
		for i := openingBracketIndex - 1; i >= start; i-- {
			switch runes[i] {
			case '-':
				reverseNextLevel = true
				break checkMinusSign
			case '+':
				break checkMinusSign
			}
		}
	}

	if reverse {
		endIndex := end
		if openingBracketIndex > 0 {
			endIndex = openingBracketIndex - 1
		}
		for i := start; i <= endIndex; i++ {
			switch runes[i] {
			case '+':
				runes[i] = '-'
			case '-':
				runes[i] = '+'
			}
		}
	}

	if openingBracketIndex < 0 {
		return
	}
	runes[openingBracketIndex] = ' '
	runes[closingBracketIndex] = ' '
	if reverseNextLevel {
		simpifyExp(runes, openingBracketIndex+1, closingBracketIndex-1, !reverse)
	} else {
		simpifyExp(runes, openingBracketIndex+1, closingBracketIndex-1, reverse)
	}
	simpifyExp(runes, closingBracketIndex+1, end, reverse)
}
