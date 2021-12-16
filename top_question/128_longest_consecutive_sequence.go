package top_question

/// Longest Consecutive Sequence
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

func LongestConsecutive(nums []int) int {
	numMap := map[int]struct{}{}
	for _, num := range nums {
		numMap[num] = struct{}{}
	}
	r := 0
	for num := range numMap {
		if _, ok := numMap[num-1]; !ok {
			curr, count := num, 1
			for _, ok2 := numMap[curr+1]; ok2; {
				curr++
				count++
				_, ok2 = numMap[curr+1]
			}
			if count > r {
				r = count
			}
		}
	}
	return r
}
