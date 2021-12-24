package top_question

import "sort"

/// Merge Intervals
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.

func Merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{intervals[0]}
	for i := 1; i < n; i++ {
		lastAns := ans[len(ans)-1]
		hasNewAns := intervals[i][0] > lastAns[1]
		var start, end int
		if lastAns[0] > intervals[i][0] {
			start = intervals[i][0]
		} else {
			start = lastAns[0]
		}
		if lastAns[1] < intervals[i][1] {
			end = intervals[i][1]
		} else {
			end = lastAns[1]
		}
		if hasNewAns {
			ans = append(ans, intervals[i])
		} else {
			ans[len(ans)-1] = []int{start, end}
		}
	}
	return ans
}
