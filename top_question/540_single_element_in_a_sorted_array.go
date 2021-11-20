package top_question

/// Single Element in a Sorted Array
// You are given a sorted array consisting of only integers where every element appears exactly twice,
// except for one element which appears exactly once.
// Return the single element that appears only once.
// Your solution must run in O(log n) time and O(1) space.
// Input: nums = [1,1,2,3,3,4,4,8,8]
// Output: 2
// Input: nums = [3,3,7,7,10,11,11]
// Output: 10

func SingleNonDuplicate(nums []int) int {
	high := len(nums) - 1
	if high == 0 {
		return nums[0]
	}
	low := 0
	var mid int
	for low < high {
		mid = (low + high) / 2
		if mid%2 == 0 {
			if nums[mid] < nums[mid+1] {
				high = mid
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] == nums[mid-1] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}
	return nums[low]
}
