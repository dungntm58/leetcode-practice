package top_question

/// Add Binary
// Given two binary strings a and b, return their sum as a binary string.
// Input: a = "11", b = "1"
// Output: "100"
// Input: a = "1010", b = "1011"
// Output: "10101"

func AddBinary(a string, b string) string {
	var bytes []byte
	if len(a) > len(b) {
		bytes = make([]byte, len(a))
	} else {
		bytes = make([]byte, len(b))
	}
	var carry byte = 0
	for i, j, z := len(a)-1, len(b)-1, 0; i >= 0 || j >= 0; i, j, z = i-1, j-1, z+1 {
		var ai, bj byte
		if i >= 0 {
			ai = a[i] - '0'
		}
		if j >= 0 {
			bj = b[j] - '0'
		}
		sum := ai + bj + carry
		carry = sum / 2
		bytes[z] = byte(sum%2 + '0')
	}
	if carry > 0 {
		bytes = append(bytes, byte(carry+'0'))
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
