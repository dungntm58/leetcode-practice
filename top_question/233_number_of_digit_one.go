package top_question

import (
	"strconv"
)

/// Number of Digit One
// Given an integer n, count the total number of digit 1 appearing in
// all non-negative integers less than or equal to n.
// Input: n = 13
// Output: 6
// Input: n = 0
// Output: 0

func CountDigitOne(n int) int {
	str := strconv.Itoa(n)
	ans := 0
	power := 1
	for i := 0; i < len(str); i++ {
		nextPower := power * 10
		roundDown := n - n%nextPower
		roundUp := roundDown + nextPower
		nextRight := n % power

		digit := str[len(str)-i-1] - '0'
		if digit < 1 {
			ans += roundDown / 10
		} else if digit == 1 {
			ans += roundDown/10 + nextRight + 1
		} else {
			ans += roundUp / 10
		}
		power = nextPower
	}
	return ans
}
