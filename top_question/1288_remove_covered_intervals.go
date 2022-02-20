package top_question

import "sort"

/// Remove Covered Intervals
// Given an array intervals where intervals[i] = [li, ri] represent the interval [li, ri),
// remove all intervals that are covered by another interval in the list.
// The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d.
// Return the number of remaining intervals.
// Input: intervals = [[1,4],[3,6],[2,8]]
// Output: 2
// Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
// Input: intervals = [[1,4],[2,3]]
// Output: 1

func RemoveCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return false
	})
	res := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][0] && intervals[i][1] <= intervals[i-1][1] {
			intervals[i] = intervals[i-1]
		} else {
			res++
		}
	}
	return res
}
