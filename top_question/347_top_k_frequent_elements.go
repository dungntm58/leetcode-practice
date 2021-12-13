package top_question

import "container/heap"

/// Top K Frequent Elements
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Input: nums = [1], k = 1
// Output: [1]

func TopKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}

	h := NewIntPriorityHeap(false)
	for key, value := range freq {
		heap.Push(h, Item{Val: key, Prority: value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	r := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		r[i] = h.Pop().(Item).Val
	}

	return r
}
