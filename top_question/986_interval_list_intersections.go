package top_question

/// Interval List Intersections
// You are given two lists of closed intervals, firstList and secondList,
// where firstList[i] = [starti, endi] and secondList[j] = [startj, endj].
// Each list of intervals is pairwise disjoint and in sorted order.
// Return the intersection of these two interval lists.
// A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.
// The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval.
// For example, the intersection of [1, 3] and [2, 4] is [2, 3].
// Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
// Input: firstList = [[1,3],[5,9]], secondList = []
// Output: []
// Input: firstList = [], secondList = [[4,8],[10,12]]
// Output: []
// Input: firstList = [[1,7]], secondList = [[3,10]]
// Output: [[3,7]]

func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	n1 := len(firstList)
	if n1 == 0 {
		return [][]int{}
	}
	n2 := len(secondList)
	if n2 == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	i := 0
	j := 0
	for i < n1 && j < n2 {
		increaseFirst := false
		var start, end int
		if firstList[i][0] > secondList[j][0] {
			start = firstList[i][0]
		} else {
			start = secondList[j][0]
		}
		if firstList[i][1] < secondList[j][1] {
			end = firstList[i][1]
			increaseFirst = true
		} else {
			end = secondList[j][1]
		}
		if start <= end {
			ans = append(ans, []int{start, end})
		}
		if increaseFirst {
			i++
		} else {
			j++
		}
	}
	return ans
}
