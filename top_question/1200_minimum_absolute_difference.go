package top_question

import "sort"

/// Minimum Absolute Difference
// Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.
// Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
// - a, b are from arr
// - a < b
// - b - a equals to the minimum absolute difference of any two elements in arr
// Input: arr = [4,2,1,3]
// Output: [[1,2],[2,3],[3,4]]
// Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
// Input: arr = [1,3,6,10,15]
// Output: [[1,3]]
// Input: arr = [3,8,-10,23,19,-4,-14,27]
// Output: [[-14,-10],[19,23],[23,27]]

func MinimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	r := [][]int{{arr[0], arr[1]}}
	minDiff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			r = [][]int{{arr[i], arr[i+1]}}
			minDiff = diff
		} else if diff == minDiff {
			r = append(r, []int{arr[i], arr[i+1]})
		}
	}
	return r
}
