package top_question

/// Nth Magical Number
// A positive integer is magical if it is divisible by either a or b.
// Given the three integers n, a, and b, return the nth magical number.
// Since the answer may be very large, return it modulo 109 + 7
// Input: n = 1, a = 2, b = 3
// Output: 2
// Input: n = 4, a = 2, b = 3
// Output: 6
// Input: n = 5, a = 2, b = 4
// Output: 10
// Input: n = 3, a = 6, b = 4
// Output: 8

func NthMagicalNumber(n int, a int, b int) int {
	mod := 1_000_000_007
	low, lcd := 0, a*b/gcd(a, b)
	var high int
	if a > b {
		high = n * b
	} else {
		high = n * a
	}
	for low < high {
		mid := (low + high) / 2
		if mid/a+mid/b-mid/lcd < n {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low % mod
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	var shift int
	for shift = 0; (a|b)&1 == 0; shift++ {
		a >>= 1
		b >>= 1
	}
	for a&1 == 0 {
		a >>= 1
	}
	for ok := true; ok; ok = b != 0 {
		for b&1 == 0 {
			b >>= 1
		}
		if a > b {
			a, b = b, a
		}
		b -= a
	}
	return a << shift
}
