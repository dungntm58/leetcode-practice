package top_question

/// House Robber II
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed.
// All houses at this place are arranged in a circle.
// That means the first house is the neighbor of the last one.
// Meanwhile, adjacent houses have a security system connected,
// and it will automatically contact the police if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.
// Input: nums = [2,3,2]
// Output: 3
// Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
// Input: nums = [1,2,3]
// Output: 3

func Rob_II(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	robFirstPre1, robFirstPre2 := nums[0], nums[0]
	if nums[1] > robFirstPre2 {
		robFirstPre2 = nums[1]
	}
	notRobFirstPre1, notRobFirstPre2 := nums[1], nums[1]
	if n >= 3 && nums[2] > notRobFirstPre2 {
		notRobFirstPre2 = nums[2]
	}
	for i := 2; i < n; i++ {
		if i < n-1 {
			robFirstPre1 += nums[i]
			r := robFirstPre1
			if r < robFirstPre2 {
				r = robFirstPre2
			}
			robFirstPre1, robFirstPre2 = robFirstPre2, r
		}
		if i > 2 {
			notRobFirstPre1 += nums[i]
			r := notRobFirstPre1
			if r < notRobFirstPre2 {
				r = notRobFirstPre2
			}
			notRobFirstPre1, notRobFirstPre2 = notRobFirstPre2, r
		}
	}
	if robFirstPre2 > notRobFirstPre2 {
		return robFirstPre2
	}
	return notRobFirstPre2
}
