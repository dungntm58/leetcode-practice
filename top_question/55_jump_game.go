package top_question

/// Jump Game
// You are given an integer array nums. You are initially positioned at the array's first index,
// and each element in the array represents your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what.
// Its maximum jump length is 0, which makes it impossible to reach the last index.

func CanJump(nums []int) bool {
	cMax := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		cMax -= 1
		if cMax < nums[i] {
			cMax = nums[i]
		}
		if cMax == 0 {
			return false
		}
	}
	return true
}
