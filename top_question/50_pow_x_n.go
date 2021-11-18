package top_question

/// Pow(x, n)
// Implement pow(x, n), which calculates x raised to the power n (i.e., xˆn).
// Input: x = 2.00000, n = 10
// Output: 1024.00000
// Input: x = 2.10000, n = 3
// Output: 9.26100
// Input: x = 2.00000, n = -2
// Output: 0.25000
// Explanation: 2ˆ-2 = 1/2ˆ2 = 1/4 = 0.25

func MyPow(x float64, n int) float64 {
	if x == 1 {
		return x
	}
	var absN int
	var nSign = true
	if n >= 0 {
		absN = n
	} else {
		absN = -n
		nSign = false
	}
	result := recursiveMultiply(x, absN)
	if nSign {
		return result
	} else {
		return 1 / result
	}
}

func recursiveMultiply(x float64, n int) float64 {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	default:
		r := recursiveMultiply(x, n/2)
		if n%2 == 1 {
			return r * r * x
		} else {
			return r * r
		}
	}
}
