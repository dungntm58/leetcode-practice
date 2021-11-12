package top_question

/// Multiply Strings
// Given two non-negative integers num1 and num2 represented as strings,
// return the product of num1 and num2, also represented as a string.
// Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
// Input: num1 = "2", num2 = "3"
// Output: "6"
// Input: num1 = "123", num2 = "456"
// Output: "56088"

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Len := len(num1)
	num2Len := len(num2)
	n1 := []rune(num1)
	n2 := []rune(num2)

	reverse(n1)
	reverse(n2)

	resultLength := num1Len*num2Len + 1
	result := make([]rune, resultLength)
	for i := 0; i < num1Len; i++ {
		d1 := n1[i] - '0'
		for j := 0; j < num2Len; j++ {
			carry := result[i+j]
			r := d1*(n2[j]-'0') + carry
			result[i+j] = r % 10
			result[i+j+1] += r / 10
		}
	}
	for result[resultLength-1] <= 0 {
		resultLength--
	}
	result = result[:resultLength]
	for i := 0; i < resultLength; i++ {
		result[i] += '0'
	}
	reverse(result)
	return string(result)
}

func reverse(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}
