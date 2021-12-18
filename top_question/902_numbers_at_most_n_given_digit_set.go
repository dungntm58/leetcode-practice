package top_question

import (
	"strconv"
)

/// Numbers At Most N Given Digit Set
// Given an array of digits which is sorted in non-decreasing order.
// You can write numbers using each digits[i] as many times as we want.
// For example, if digits = ['1','3','5'], we may write numbers such as '13', '551', and '1351315'.
// Return the number of positive integers that can be generated that are less than or equal to a given integer n.
// Input: digits = ["1","3","5","7"], n = 100
// Output: 20
// Explanation:
// The 20 numbers that can be written are:
// 1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
// Input: digits = ["1","4","9"], n = 1000000000
// Output: 29523
// Explanation:
// We can write 3 one digit numbers, 9 two digit numbers, 27 three digit numbers,
// 81 four digit numbers, 243 five digit numbers, 729 six digit numbers,
// 2187 seven digit numbers, 6561 eight digit numbers, and 19683 nine digit numbers.
// In total, this is 29523 integers that can be written using the digits array.
// Input: digits = ["7"], n = 8
// Output: 1

func AtMostNGivenDigitSet(digits []string, n int) int {
	nStr := strconv.Itoa(n)
	digitsCount := len(digits)
	r, power := 0, digitsCount
	for i := 1; i < len(nStr); i++ {
		r += power
		power *= digitsCount
	}
	power /= digitsCount
	mostSimilarNumberCount := 1
	for _, d := range nStr {
		maxDigitMostFirstIndex := len(digits) - 1
		for maxDigitMostFirstIndex >= 0 && digits[maxDigitMostFirstIndex][0] > byte(d) {
			maxDigitMostFirstIndex--
		}
		if maxDigitMostFirstIndex < 0 {
			mostSimilarNumberCount = 0
			break
		}
		if digits[maxDigitMostFirstIndex][0] == byte(d) {
			r += maxDigitMostFirstIndex * power
		} else {
			mostSimilarNumberCount = 0
			r += (maxDigitMostFirstIndex + 1) * power
			break
		}
		power /= digitsCount
	}
	r += mostSimilarNumberCount
	return r
}
