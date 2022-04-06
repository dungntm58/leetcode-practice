package top_question

import "sort"

/// 3Sum With Multiplicity
// Given an integer array arr, and an integer target, return the number of tuples i, j, k
// such that i < j < k and arr[i] + arr[j] + arr[k] == target.
// As the answer can be very large, return it modulo 109 + 7.
// Input: arr = [1,1,2,2,3,3,4,4,5,5], target = 8
// Output: 20
// Explanation:
// Enumerating by the values (arr[i], arr[j], arr[k]):
// (1, 2, 5) occurs 8 times;
// (1, 3, 4) occurs 8 times;
// (2, 2, 4) occurs 2 times;
// (2, 3, 3) occurs 2 times.
// Input: arr = [1,1,2,2,2,2], target = 5
// Output: 12
// Explanation:
// arr[i] = 1, arr[j] = arr[k] = 2 occurs 12 times:
// We choose one 1 from [1,1] in 2 ways,
// and two 2s from [2,2,2,2] in 6 ways.

func ThreeSumMulti(arr []int, target int) int {
	MOD := 1_000_000_007
	ans := 0
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		T := target - arr[i]
		j, k := i+1, len(arr)-1
		for j < k {
			if arr[j]+arr[k] < T {
				j++
			} else if arr[j]+arr[k] > T {
				k--
			} else if arr[j] != arr[k] {
				left, right := 1, 1
				for j+1 < k && arr[j] == arr[j+1] {
					left++
					j++
				}
				for k-1 > j && arr[k] == arr[k-1] {
					right++
					k--
				}
				ans += left * right
				ans %= MOD
				j++
				k--
			} else {
				ans += (k - j + 1) * (k - j) / 2
				ans %= MOD
				break
			}
		}
	}
	return ans
}
