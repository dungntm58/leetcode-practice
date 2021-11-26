package top_question

import "fmt"

/// Search Insert Position
// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.
// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Input: nums = [1,3,5,6], target = 7
// Output: 4
// Input: nums = [1,3,5,6], target = 0
// Output: 0
// Input: nums = [1], target = 0
// Output: 0

func SearchInsert(nums []int, target int) int {
	low, mid, high := 0, 0, len(nums)-1
	for low < high {
		mid = (low + high) / 2
		fmt.Println("mid =", mid, "low =", low, "high =", high)
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid
		} else {
			return mid
		}
	}
	if nums[low] < target {
		return low + 1
	}
	return low
}
