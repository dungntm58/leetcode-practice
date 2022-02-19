package top_question

import (
	"container/heap"
)

/// Minimize Deviation in Array
// You are given an array nums of n positive integers.
// You can perform two types of operations on any element of the array any number of times:
// - If the element is even, divide it by 2.
//   For example, if the array is [1,2,3,4], then you can do this operation on the last element,
//   and the array will be [1,2,3,2].
// - If the element is odd, multiply it by 2.
//   For example, if the array is [1,2,3,4], then you can do this operation on the first element,
//   and the array will be [2,2,3,4].
// The deviation of the array is the maximum difference between any two elements in the array.
// Return the minimum deviation the array can have after performing some number of operations.
// Input: nums = [1,2,3,4]
// Output: 1
// Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.
// Input: nums = [4,1,5,20,3]
// Output: 3
// Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.
// Input: nums = [2,10,8]
// Output: 3

func MinimumDeviation(nums []int) int {
	intHeap := NewIntHeap(true)
	min, max := 1_000_000_000, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			nums[i] *= 2
		}
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
		heap.Push(intHeap, nums[i])
	}
	res := max - min
	for intHeap.Peek().(int)%2 == 0 {
		top := heap.Pop(intHeap).(int)
		if res > top-min {
			res = top - min
		}
		top /= 2
		if min > top {
			min = top
		}
		heap.Push(intHeap, top)
	}
	if res > intHeap.Peek().(int)-min {
		res = intHeap.Peek().(int) - min
	}
	return res
}
