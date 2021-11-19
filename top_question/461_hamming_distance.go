package top_question

import "strconv"

/// Hamming Distance
// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
// Given two integers x and y, return the Hamming distance between them.
// Input: x = 1, y = 4
// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑
// The above arrows point to positions where the corresponding bits are different.
// Input: x = 3, y = 1
// Output: 1

func HammingDistance(x int, y int) int {
	r := 0
	if x > y {
		x, y = y, x
	}
	xArr := []rune(strconv.FormatInt(int64(x), 2))
	yArr := []rune(strconv.FormatInt(int64(y), 2))
	reverse(xArr)
	reverse(yArr)

	xn := len(xArr)
	for i, ye := range yArr {
		if i < xn {
			if ye != xArr[i] {
				r++
			}
		} else if ye == '1' {
			r++
		}
	}
	return r
}

// func hammingDistance(x int, y int) int {
// 	diff:=x^y
// 	ans:=0
// 	for diff>0{
// 			diff=diff&(diff-1)
// 			ans++
// 	}
// 	return ans
// }
