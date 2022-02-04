package top_question

/// Contiguous Array
// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.
// Input: nums = [0,1]
// Output: 2
// Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
// Input: nums = [0,1,0]
// Output: 2
// Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

func FindMaxLength(nums []int) int {
	indexCountMap := map[int]int{}
	indexCountMap[0] = -1
	maxLength, count := 0, 0
	for i, num := range nums {
		if num == 0 {
			count--
		} else {
			count++
		}
		if _, ok := indexCountMap[count]; ok {
			if maxLength < i-indexCountMap[count] {
				maxLength = i - indexCountMap[count]
			}
		} else {
			indexCountMap[count] = i
		}
	}
	return maxLength
}

// func FindMaxLength(nums []int) int {
// 	arr := make([]int, 2*len(nums)+1)
// 	for i, _ := range arr {
// 		arr[i] = -2
// 	}
// 	arr[len(nums)] = -1
// 	max, count := 0, 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			count--
// 		} else {
// 			count++
// 		}
// 		if arr[count+len(nums)] >= -1 {
// 			if max < i-arr[count+len(nums)] {
// 				max = i - arr[count+len(nums)]
// 			}
// 		} else {
// 			arr[count+len(nums)] = i
// 		}
// 	}
// 	return max
// }
