package top_question

/// Valid Mountain Array
// Given an array of integers arr, return true if and only if it is a valid mountain array.
// Recall that arr is a mountain array if and only if:
// - arr.length >= 3
// - There exists some i with 0 < i < arr.length - 1 such that:
//     arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//     arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
// Input: arr = [2,1]
// Output: false
// Input: arr = [3,5,5]
// Output: false
// Input: arr = [0,3,2,1]
// Output: true

func ValidMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	topIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			topIndex = i - 1
			break
		} else if arr[i] == arr[i-1] {
			return false
		}
	}
	if topIndex == 0 {
		return false
	}
	for i := topIndex + 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}
	return true
}
