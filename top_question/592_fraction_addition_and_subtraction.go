package top_question

import "fmt"

// Given a string expression representing an expression of fraction addition and subtraction, return the calculation result in string format.
// The final result should be an irreducible fraction. If your final result is an integer, change it to the format of a fraction that has a denominator 1. So in this case, 2 should be converted to 2/1.

// Example 1:

// Input: expression = "-1/2+1/2"
// Output: "0/1"
// Example 2:

// Input: expression = "-1/2+1/2+1/3"
// Output: "1/3"
// Example 3:

// Input: expression = "1/3-1/2"
// Output: "-1/6"

// Constraints:
// - The input string only contains '0' to '9', '/', '+' and '-'. So does the output.
// - Each fraction (input and output) has the format ±numerator/denominator. If the first input fraction or the output is positive, then '+' will be omitted.
// - The input only contains valid irreducible fractions, where the numerator and denominator of each fraction will always be in the range [1, 10]. If the denominator is 1, it means this fraction is actually an integer in a fraction format defined above.
// - The number of given fractions will be in the range [1, 10].
// - The numerator and denominator of the final result are guaranteed to be valid and in the range of 32-bit int.

func FractionAddition(expression string) string {
	var isNegative = false
	var sNumerator, sDenominator = 0, 1
	var numerator, denominator = 0, 0
	var isNumerator = true
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '-':
			isNegative = true
			isNumerator = true
			sNumerator, sDenominator = addFraction(sNumerator, sDenominator, numerator, denominator)
			numerator, denominator = 0, 0
		case '+':
			isNegative = false
			isNumerator = true
			sNumerator, sDenominator = addFraction(sNumerator, sDenominator, numerator, denominator)
			numerator, denominator = 0, 0
		case '/':
			isNumerator = false
			if isNegative {
				numerator = -numerator
			}
		default:
			digit := expression[i] - '0'
			if isNumerator {
				numerator *= 10
				numerator += int(digit)
			} else {
				denominator *= 10
				denominator += int(digit)
			}
		}
	}
	sNumerator, sDenominator = addFraction(sNumerator, sDenominator, numerator, denominator)
	return fmt.Sprintf("%d/%d", sNumerator, sDenominator)
}

func addFraction(lNumerator, lDenominator, rNumerator, rDenominator int) (int, int) {
	if lNumerator == 0 {
		if rNumerator == 0 {
			return 0, 1
		}
		return rNumerator, max(rDenominator, 1)
	} else if rNumerator == 0 {
		return lNumerator, max(lDenominator, 1)
	}
	var numerator, denominator int
	if lDenominator == rDenominator {
		numerator, denominator = lNumerator+rNumerator, lDenominator
	} else {
		numerator, denominator = lNumerator*rDenominator+lDenominator*rNumerator, lDenominator*rDenominator
	}
	if numerator == 0 {
		return 0, 1
	}
	var _gcd = 1
	if numerator > 0 {
		_gcd = gcd(numerator, denominator)
	} else {
		_gcd = gcd(-numerator, denominator)
	}
	return numerator / _gcd, denominator / _gcd
}
