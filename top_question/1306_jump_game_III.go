package top_question

/// Jump Game III
// Given an array of non-negative integers arr, you are initially positioned at start index of the array.
// When you are at index i, you can jump to i + arr[i] or i - arr[i], check if you can reach to any index with value 0.
// Notice that you can not jump outside of the array at any time.
// Input: arr = [4,2,3,0,3,1,2], start = 5
// Output: true
// Explanation:
// All possible ways to reach at index 3 with value 0 are:
// index 5 -> index 4 -> index 1 -> index 3
// index 5 -> index 6 -> index 4 -> index 1 -> index 3
// Input: arr = [4,2,3,0,3,1,2], start = 0
// Output: true
// Explanation:
// One possible way to reach at index 3 with value 0 is:
// index 0 -> index 4 -> index 1 -> index 3
// Input: arr = [3,0,2,1,2], start = 2
// Output: false
// Explanation: There is no way to reach at index 1 with value 0.

func CanReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	return checkReach(arr, start, visited)
}

func checkReach(arr []int, current int, visited []bool) bool {
	if current < 0 || current >= len(arr) {
		return false
	}
	if arr[current] == 0 {
		return true
	}
	if visited[current] {
		return false
	}
	visited[current] = true
	if current < arr[current] {
		return checkReach(arr, current+arr[current], visited)
	}
	if current+arr[current] > len(arr)-1 {
		return checkReach(arr, current-arr[current], visited)
	}
	return checkReach(arr, current+arr[current], visited) || checkReach(arr, current-arr[current], visited)
}
