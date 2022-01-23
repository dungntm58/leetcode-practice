package top_question

import (
	"strconv"
)

/// Sequential Digits
// An integer has sequential digits if and only if each digit in the number is one more than the previous digit.
// Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.
// Input: low = 100, high = 300
// Output: [123,234]
// Input: low = 1000, high = 13000
// Output: [1234,2345,3456,4567,5678,6789,12345]

func SequentialDigits(low int, high int) []int {
	variance, digit, index := 1, 0, -1
	lowStr := strconv.Itoa(low)
	for i := 1; i < len(lowStr); i++ {
		variance *= 10
		variance++
	}
	for i := 1; i < len(lowStr); i++ {
		if int(lowStr[i])-int(lowStr[i-1]) > 1 {
			digit = int(lowStr[i-1]-'0') + 1
			index = i - 1
			break
		}
		if lowStr[i] <= lowStr[i-1] {
			digit = int(lowStr[i-1] - '0')
			index = i - 1
			break
		}
	}
	start, digitCount := 0, len(lowStr)
	if index == -1 {
		start = low
	} else {
		digit += len(lowStr) - index - 1
		if digit <= 9 {
			for i := digit - len(lowStr) + 1; i <= digit; i++ {
				start = start*10 + i
			}
		} else {
			digitCount++
			variance = variance*10 + 1
			digit = len(lowStr) + 1
			for i := 1; i <= digit; i++ {
				start = start*10 + i
			}
		}
	}
	res := []int{}
	for start <= high {
		res = append(res, start)
		if start%10 == 9 {
			digitCount++
			start = 0
			for i := 1; i <= digitCount; i++ {
				start = start*10 + i
			}
			variance = variance*10 + 1
		} else {
			start += variance
		}
	}
	return res
}
