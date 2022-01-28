package contest

// 2138
/// Divide a String Into Groups of Size k
// A string s can be partitioned into groups of size k using the following procedure:
// - The first group consists of the first k characters of the string,
// the second group consists of the next k characters of the string, and so on.
// Each character can be a part of exactly one group.
// - For the last group, if the string does not have k characters remaining,
// a character fill is used to complete the group.
// Note that the partition is done so that after removing the fill character
// from the last group (if it exists) and concatenating all the groups in order, the resultant string should be s.
// Given the string s, the size of each group k and the character fill,
// return a string array denoting the composition of every group s has been divided into, using the above procedure.
// Input: s = "abcdefghi", k = 3, fill = "x"
// Output: ["abc","def","ghi"]
// Explanation:
// The first 3 characters "abc" form the first group.
// The next 3 characters "def" form the second group.
// The last 3 characters "ghi" form the third group.
// Since all groups can be completely filled by characters from the string, we do not need to use fill.
// Thus, the groups formed are "abc", "def", and "ghi".
// Input: s = "abcdefghij", k = 3, fill = "x"
// Output: ["abc","def","ghi","jxx"]
// Explanation:
// Similar to the previous example, we are forming the first three groups "abc", "def", and "ghi".
// For the last group, we can only use the character 'j' from the string. To complete this group, we add 'x' twice.
// Thus, the 4 groups formed are "abc", "def", "ghi", and "jxx".

func DivideString(s string, k int, fill byte) []string {
	count := 0
	res := []string{}
	for count+k <= len(s) {
		res = append(res, s[count:count+k])
		count += k
	}
	if count < len(s) {
		lastRunes := []byte(s[count:])
		for i := len(s) - count; i < k; i++ {
			lastRunes = append(lastRunes, fill)
		}
		res = append(res, string(lastRunes))
	}
	return res
}

// 2139
/// Minimum Moves to Reach Target Score
// You are playing a game with integers. You start with the integer 1 and you want to reach the integer target.
// In one move, you can either:
// - Increment the current integer by one (i.e., x = x + 1).
// - Double the current integer (i.e., x = 2 * x).
// You can use the increment operation any number of times,
// however, you can only use the double operation at most maxDoubles times.
// Given the two integers target and maxDoubles, return the minimum number of moves needed to reach target starting with 1.
// Input: target = 5, maxDoubles = 0
// Output: 4
// Explanation: Keep incrementing by 1 until you reach target.
// Input: target = 19, maxDoubles = 2
// Output: 7
// Explanation: Initially, x = 1
// Increment 3 times so x = 4
// Double once so x = 8
// Increment once so x = 9
// Double again so x = 18
// Increment once so x = 19
// Input: target = 10, maxDoubles = 4
// Output: 4
// Explanation: Initially, x = 1
// Increment once so x = 2
// Double once so x = 4
// Increment once so x = 5
// Double again so x = 10

func MinMoves(target int, maxDoubles int) int {
	count := 0
	for maxDoubles > 0 && target > 1 {
		if target%2 == 0 {
			target /= 2
			maxDoubles--
		} else {
			target--
		}
		count++
	}
	return count + target - 1
}
