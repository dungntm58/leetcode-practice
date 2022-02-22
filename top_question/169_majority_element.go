package top_question

/// Majority Element
// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
// Input: nums = [3,2,3]
// Output: 3
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

func MajorityElement(nums []int) int {
	res := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if res == v {
			count++
		} else {
			count--
		}
	}
	return res
}

// func MajorityElement(nums []int) int {
// 	numMap := make(map[int]int)
// 	for _, num := range nums {
// 		numMap[num]++
// 	}
// 	for k, v := range numMap {
// 		if v > len(nums)/2 {
// 			return k
// 		}
// 	}
// 	return 0
// }
