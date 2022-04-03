package top_question

/// Search in Rotated Sorted Array II
// There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
// Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
// such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
// Given the array nums after the rotation and an integer target, return true
// if target is in nums, or false if it is not in nums.
// You must decrease the overall operation steps as much as possible.
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true
// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false

func Search(nums []int, target int) bool {
	breakPoint := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			breakPoint = i
			break
		}
	}
	leftIndex := binarySearchInRange(nums, target, 0, breakPoint-1)
	if leftIndex == -1 {
		return binarySearchInRange(nums, target, breakPoint, len(nums)-1) != -1
	}
	return true
}

func binarySearchInRange(arr []int, target int, start, end int) int {
	for start <= end {
		mid := (end + start) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
