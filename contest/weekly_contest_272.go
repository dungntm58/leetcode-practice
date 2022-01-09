package contest

// 2108
/// Find First Palindromic String in the Array
// Given an array of strings words, return the first palindromic string in the array.
// If there is no such string, return an empty string "".
// A string is palindromic if it reads the same forward and backward.
// Input: words = ["abc","car","ada","racecar","cool"]
// Output: "ada"
// Explanation: The first string that is palindromic is "ada".
// Note that "racecar" is also palindromic, but it is not the first.
// Input: words = ["notapalindrome","racecar"]
// Output: "racecar"
// Explanation: The first and only string that is palindromic is "racecar".
// Input: words = ["def","ghi"]
// Output: ""
// Explanation: There are no palindromic strings, so the empty string is returned.

func FirstPalindrome(words []string) string {
LOOP:
	for _, w := range words {
		for i := 0; i < len(w)/2+1; i++ {
			if w[i] != w[len(w)-1-i] {
				continue LOOP
			}
		}
		return w
	}
	return ""
}

// func FirstPalindrome(words []string) string {
// 	for _, w := range words {
// 		l := len(w)
// 		lC, rC := l/2, l/2
// 		if l%2 == 0 {
// 			rC++
// 		}
// 		for lC >= 0 && rC < l && w[lC] == w[rC] {
// 			lC--
// 			rC++
// 		}
// 		if lC == -1 {
// 			return w
// 		}
// 	}
// 	return ""
// }

// 2109
/// Adding Spaces to a String
// Input: s = "LeetcodeHelpsMeLearn", spaces = [8,13,15]
// Output: "Leetcode Helps Me Learn"
// Explanation:
// The indices 8, 13, and 15 correspond to the underlined characters in "LeetcodeHelpsMeLearn".
// We then place spaces before those characters.
// Input: s = "icodeinpython", spaces = [1,5,7,9]
// Output: "i code in py thon"
// Explanation:
// The indices 1, 5, 7, and 9 correspond to the underlined characters in "icodeinpython".
// We then place spaces before those characters.
// Input: s = "spacing", spaces = [0,1,2,3,4,5,6]
// Output: " s p a c i n g"
// Explanation:
// We are also able to place spaces before the first character of the string.

func AddSpaces(s string, spaces []int) string {
	r := make([]byte, len(s)+len(spaces))
	for i, s := range spaces {
		r[i+s] = ' '
	}
	i := 0
	for j := 0; j < len(s); j++ {
		if r[i] == ' ' {
			i++
		}
		r[i] = s[j]
		i++
	}
	return string(r)
}

// 2110
/// Number of Smooth Descent Periods of a Stock
// You are given an integer array prices representing the daily price history of a stock,
// where prices[i] is the stock price on the ith day.
// A smooth descent period of a stock consists of one or more contiguous days
// such that the price on each day is lower than the price on the preceding day by exactly 1.
// The first day of the period is exempted from this rule.
// Return the number of smooth descent periods.
// Input: prices = [3,2,1,4]
// Output: 7
// Explanation: There are 7 smooth descent periods:
// [3], [2], [1], [4], [3,2], [2,1], and [3,2,1]
// Note that a period with one day is a smooth descent period by the definition.
// Input: prices = [8,6,7,7]
// Output: 4
// Explanation: There are 4 smooth descent periods: [8], [6], [7], and [7]
// Note that [8,6] is not a smooth descent period as 8 - 6 ≠ 1.
// Input: prices = [1]
// Output: 1
// Explanation: There is 1 smooth descent period: [1]

func GetDescentPeriods(prices []int) int64 {
	return 0
}

// 2111
/// Minimum Operations to Make the Array K-Increasing
// You are given a 0-indexed array arr consisting of n positive integers, and a positive integer k.
// The array arr is called K-increasing if arr[i-k] <= arr[i] holds for every index i, where k <= i <= n-1.
// For example, arr = [4, 1, 5, 2, 6, 2] is K-increasing for k = 2 because:
// arr[0] <= arr[2] (4 <= 5)
// arr[1] <= arr[3] (1 <= 2)
// arr[2] <= arr[4] (5 <= 6)
// arr[3] <= arr[5] (2 <= 2)
// However, the same arr is not K-increasing for k = 1 (because arr[0] > arr[1]) or k = 3 (because arr[0] > arr[3]).
// In one operation, you can choose an index i and change arr[i] into any positive integer.
// Return the minimum number of operations required to make the array K-increasing for the given k.
// Input: arr = [5,4,3,2,1], k = 1
// Output: 4
// Explanation:
// For k = 1, the resultant array has to be non-decreasing.
// Some of the K-increasing arrays that can be formed are [5,6,7,8,9], [1,1,1,1,1], [2,2,3,4,4]. All of them require 4 operations.
// It is suboptimal to change the array to, for example, [6,7,8,9,10] because it would take 5 operations.
// It can be shown that we cannot make the array K-increasing in less than 4 operations.
// Input: arr = [4,1,5,2,6,2], k = 2
// Output: 0
// Explanation:
// This is the same example as the one in the problem description.
// Here, for every index i where 2 <= i <= 5, arr[i-2] <= arr[i].
// Since the given array is already K-increasing, we do not need to perform any operations.
// Input: arr = [4,1,5,2,6,2], k = 3
// Output: 2
// Explanation:
// Indices 3 and 5 are the only ones not satisfying arr[i-3] <= arr[i] for 3 <= i <= 5.
// One of the ways we can make the array K-increasing is by changing arr[3] to 4 and arr[5] to 5.
// The array will now be [4,1,5,4,6,5].
// Note that there can be other ways to make the array K-increasing, but none of them require less than 2 operations.

func KIncreasing(arr []int, k int) int {
	return 0
}
