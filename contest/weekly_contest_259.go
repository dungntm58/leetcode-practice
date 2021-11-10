package contest

import (
	"strings"
)

/// Final Value of Variable After Performing Operations
// There is a programming language with only four operations and one variable X:
// ++X and X++ increments the value of the variable X by 1.
// --X and X-- decrements the value of the variable X by 1.
// Initially, the value of X is 0.
// Given an array of strings operations containing a list of operations,
// return the final value of X after performing all the operations.
// Input: operations = ["--X","X++","X++"]
// Output: 1
// Explanation: The operations are performed as follows:
// Initially, X = 0.
// --X: X is decremented by 1, X =  0 - 1 = -1.
// X++: X is incremented by 1, X = -1 + 1 =  0.
// X++: X is incremented by 1, X =  0 + 1 =  1.
// Input: operations = ["++X","++X","X++"]
// Output: 3
// Explanation: The operations are performed as follows:
// Initially, X = 0.
// ++X: X is incremented by 1, X = 0 + 1 = 1.
// ++X: X is incremented by 1, X = 1 + 1 = 2.
// X++: X is incremented by 1, X = 2 + 1 = 3.
// Input: operations = ["X++","++X","--X","X--"]
// Output: 0
// Explanation: The operations are performed as follows:
// Initially, X = 0.
// X++: X is incremented by 1, X = 0 + 1 = 1.
// ++X: X is incremented by 1, X = 1 + 1 = 2.
// --X: X is decremented by 1, X = 2 - 1 = 1.
// X--: X is decremented by 1, X = 1 - 1 = 0.

func FinalValueAfterOperations(operations []string) int {
	result := 0
	for _, operation := range operations {
		if strings.HasPrefix(operation, "++") || strings.HasSuffix(operation, "++") {
			result += 1
		} else if strings.HasPrefix(operation, "--") || strings.HasSuffix(operation, "--") {
			result -= 1
		}
	}
	return result
}

/// Sum of Beauty in the Array
// You are given a 0-indexed integer array nums.
// For each index i (1 <= i <= nums.length - 2) the beauty of nums[i] equals:
// * 2, if nums[j] < nums[i] < nums[k], for all 0 <= j < i and for all i < k <= nums.length - 1.
// * 1, if nums[i - 1] < nums[i] < nums[i + 1], and the previous condition is not satisfied.
// * 0, if none of the previous conditions holds.
// Return the sum of beauty of all nums[i] where 1 <= i <= nums.length - 2.
// Input: nums = [1,2,3]
// Output: 2
// Explanation: For each index i in the range 1 <= i <= 1:
// - The beauty of nums[1] equals 2.
// Input: nums = [2,4,6,4]
// Output: 1
// Explanation: For each index i in the range 1 <= i <= 2:
// - The beauty of nums[1] equals 1.
// - The beauty of nums[2] equals 0.
// Input: nums = [3,2,1]
// Output: 0
// Explanation: For each index i in the range 1 <= i <= 1:
// - The beauty of nums[1] equals 0.

func SumOfBeauties(nums []int) int {
	length := len(nums)

	minValues := make([]int, length)
	min := nums[length-1]
	minValues[length-1] = min
	for i := length - 2; i >= 2; i-- {
		if min > nums[i] {
			min = nums[i]
		}
		minValues[i] = min
	}

	result := 0
	var leftMax int
	nextLeftMax := nums[0]

	for currI := 1; currI < length-1; currI++ {
		leftMax = nextLeftMax
		if nextLeftMax < nums[currI] {
			nextLeftMax = nums[currI]
		}
		beauty := 0
		if leftMax < nums[currI] && nums[currI] < minValues[currI+1] {
			beauty = 2
		} else if nums[currI-1] < nums[currI] && nums[currI] < nums[currI+1] {
			beauty = 1
		}
		result += beauty
	}
	return result
}

/// Detect Squares
// You are given a stream of points on the X-Y plane. Design an algorithm that:
// Adds new points from the stream into a data structure.
// Duplicate points are allowed and should be treated as different points.
// Given a query point,
// counts the number of ways to choose three points from the data structure such that the three points
// and the query point form an axis-aligned square with positive area.
// An axis-aligned square is a square whose edges are all the same length and are either parallel or perpendicular to the x-axis and y-axis.
// Implement the DetectSquares class:
// - DetectSquares() Initializes the object with an empty data structure.
// - void add(int[] point) Adds a new point point = [x, y] to the data structure.
// - int count(int[] point) Counts the number of ways to form axis-aligned squares with point point = [x, y] as described above.
// Input
// ["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
// [[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
// Output
// [null, null, null, null, 1, 0, null, 2]

// Explanation
// DetectSquares detectSquares = new DetectSquares();
// detectSquares.add([3, 10]);
// detectSquares.add([11, 2]);
// detectSquares.add([3, 2]);
// detectSquares.count([11, 10]); // return 1. You can choose:
//                                //   - The first, second, and third points
// detectSquares.count([14, 8]);  // return 0. The query point cannot form a square with any points in the data structure.
// detectSquares.add([11, 2]);    // Adding duplicate points is allowed.
// detectSquares.count([11, 10]); // return 2. You can choose:
//   															// - The first, second, and third points
//   															// - The first, third, and fourth points

type pt struct{ x, y int }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type DetectSquares struct {
	points map[pt]int
}

func Constructor() DetectSquares {
	return DetectSquares{points: map[pt]int{}}
}

func (this *DetectSquares) Add(point []int) {
	this.points[pt{x: point[0], y: point[1]}]++
}

func (this *DetectSquares) countSqrWithExistingThreePts(points []pt) int {
	occurrence := [3]int{}
	for i, p := range points {
		occurrence[i] = this.points[p]
	}
	result := 1
	for _, o := range occurrence {
		if o == 0 {
			return 0
		}
		result *= o
	}
	return result
}

func (this *DetectSquares) Count(point []int) int {
	result := 0
	sameYPts := []pt{}
	for p := range this.points {
		if p.y == point[1] && p.x != point[0] {
			sameYPts = append(sameYPts, p)
		}
	}
	for _, p := range sameYPts {
		edge := abs(p.x - point[0])
		upSideSqrPt := pt{x: p.x, y: p.y + edge}
		var downSideSqrPt *pt
		if p.y >= edge {
			downSideSqrPt = &pt{x: p.x, y: p.y - edge}
		}
		var upSideSqr int
		var downSideSqr int
		if this.points[upSideSqrPt] > 0 {
			lastPt := pt{x: point[0], y: upSideSqrPt.y}
			upSideSqr = this.countSqrWithExistingThreePts([]pt{p, upSideSqrPt, lastPt})
		}
		if downSideSqrPt != nil && this.points[*downSideSqrPt] > 0 {
			lastPt := pt{x: point[0], y: downSideSqrPt.y}
			downSideSqr = this.countSqrWithExistingThreePts([]pt{p, *downSideSqrPt, lastPt})
		}
		result += upSideSqr + downSideSqr
	}

	return result
}

/// Longest Subsequence Repeated k Times
// You are given a string s of length n, and an integer k. You are tasked to find the longest subsequence repeated k times in string s.
// A subsequence is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.
// A subsequence seq is repeated k times in the string s if seq * k is a subsequence of s,
// where seq * k represents a string constructed by concatenating seq k times.
// For example, "bba" is repeated 2 times in the string "bababcba", because the string "bbabba",
// constructed by concatenating "bba" 2 times, is a subsequence of the string "bababcba".
// Return the longest subsequence repeated k times in string s.
// If multiple such subsequences are found, return the lexicographically largest one. If there is no such subsequence, return an empty string.
// Input: s = "letsleetcode", k = 2
// Output: "let"
// Explanation: There are two longest subsequences repeated 2 times: "let" and "ete".
// "let" is the lexicographically largest one.
// Input: s = "bb", k = 2
// Output: "b"
// Explanation: The longest subsequence repeated 2 times is "b".
// Input: s = "ab", k = 2
// Output: ""
// Explanation: There is no subsequence repeated 2 times. Empty string is returned.
// Input: s = "bbabbabbbbabaababab", k = 3
// Output: "bbbb"
// Explanation: The longest subsequence "bbbb" is repeated 3 times in "bbabbabbbbabaababab".
func LongestSubsequenceRepeatedK(s string, k int) string {
	return ""
}
