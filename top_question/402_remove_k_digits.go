package top_question

/// Remove K Digits
// Given string num representing a non-negative integer num, and an integer k,
// return the smallest possible integer after removing k digits from num.
// Input: num = "1432219", k = 3
// Output: "1219"
// Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
// Input: num = "10200", k = 1
// Output: "200"
// Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
// Input: num = "10", k = 2
// Output: "0"
// Explanation: Remove all the digits from the number and it is left with nothing which is 0.

func RemoveKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	res := []byte{}
	for i := 0; i < len(num); i++ {
		digit := num[i]
		for len(res) > 0 && res[len(res)-1] > num[i] && k > 0 {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, digit)
	}
	for len(res) > 0 && k > 0 {
		res = res[:len(res)-1]
		k--
	}
	for len(res) > 0 && res[0] == '0' {
		res = res[1:]
	}
	if len(res) == 0 {
		return "0"
	}
	return string(res)
}
