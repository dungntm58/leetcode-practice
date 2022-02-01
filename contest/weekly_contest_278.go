package contest

// 2154
/// Keep Multiplying Found Values by Two
// You are given an array of integers nums.
// You are also given an integer original which is the first number that needs to be searched for in nums.
// You then do the following steps:
// - If original is found in nums, multiply it by two (i.e., set original = 2 * original).
// - Otherwise, stop the process.
// - Repeat this process with the new number as long as you keep finding the number.
// Return the final value of original.
// Input: nums = [5,3,6,1,12], original = 3
// Output: 24
// Explanation:
// - 3 is found in nums. 3 is multiplied by 2 to obtain 6.
// - 6 is found in nums. 6 is multiplied by 2 to obtain 12.
// - 12 is found in nums. 12 is multiplied by 2 to obtain 24.
// - 24 is not found in nums. Thus, 24 is returned.
// Input: nums = [2,7,9], original = 4
// Output: 4
// Explanation:
// - 4 is not found in nums. Thus, 4 is returned.

func FindFinalValue(nums []int, original int) int {
	numMap := map[int]struct{}{}
	for _, num := range nums {
		numMap[num] = struct{}{}
	}
	for {
		if _, ok := numMap[original]; ok {
			original *= 2
		} else {
			break
		}
	}
	return original
}
