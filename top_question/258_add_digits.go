package top_question

/// Add Digits
// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
// Input: num = 38
// Output: 2
// Explanation: The process is
// 38 --> 3 + 8 --> 11
// 11 --> 1 + 1 --> 2
// Since 2 has only one digit, return it.
// Input: num = 0
// Output: 0

func AddDigits(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + (num-1)%9
}

// func AddDigits(num int) int {
// 	if num < 10 {
// 		return num
// 	}
// 	sum := 0
// 	for num > 0 {
// 		sum += num % 10
// 		num /= 10
// 	}
// 	return AddDigits(sum)
// }
