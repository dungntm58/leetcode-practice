package top_question

/// Maximum XOR of Two Numbers in an Array
// Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 <= i <= j < n.
// Input: nums = [3,10,5,25,2,8]
// Output: 28
// Explanation: The maximum result is 5 XOR 25 = 28.
// Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
// Output: 127

type bitTrieNode struct {
	children [2]*bitTrieNode
}

func FindMaximumXOR(nums []int) int {
	root := &bitTrieNode{}
	for _, num := range nums {
		cur := root
		for i := 31; i >= 0; i-- {
			bit := (num >> i) & 1
			if cur.children[bit] == nil {
				cur.children[bit] = &bitTrieNode{}
			}
			cur = cur.children[bit]
		}
	}
	max := 0
	for _, num := range nums {
		cur := root
		res := 0
		for i := 31; i >= 0; i-- {
			bit := (num >> i) & 1
			if cur.children[bit^1] != nil {
				res += 1 << i
				cur = cur.children[bit^1]
			} else {
				cur = cur.children[bit]
			}
		}
		if res > max {
			max = res
		}
	}
	return max
}
