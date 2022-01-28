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

// 2140
/// Solving Questions With Brainpower
// You are given a 0-indexed 2D integer array questions where questions[i] = [pointsi, brainpoweri].
// The array describes the questions of an exam, where you have to process the questions in order
// (i.e., starting from question 0) and make a decision whether to solve or skip each question.
// Solving question i will earn you pointsi points but you will be unable to solve each of the next brainpoweri questions.
// If you skip question i, you get to make the decision on the next question.
// For example, given questions = [[3, 2], [4, 3], [4, 4], [2, 5]]:
// If question 0 is solved, you will earn 3 points but you will be unable to solve questions 1 and 2.
// If instead, question 0 is skipped and question 1 is solved,
// you will earn 4 points but you will be unable to solve questions 2 and 3.
// Return the maximum points you can earn for the exam.
// Input: questions = [[3,2],[4,3],[4,4],[2,5]]
// Output: 5
// Explanation: The maximum points can be earned by solving questions 0 and 3.
// - Solve question 0: Earn 3 points, will be unable to solve the next 2 questions
// - Unable to solve questions 1 and 2
// - Solve question 3: Earn 2 points
// Total points earned: 3 + 2 = 5. There is no other way to earn 5 or more points.
// Input: questions = [[1,1],[2,2],[3,3],[4,4],[5,5]]
// Output: 7
// Explanation: The maximum points can be earned by solving questions 1 and 4.
// - Skip question 0
// - Solve question 1: Earn 2 points, will be unable to solve the next 2 questions
// - Unable to solve questions 2 and 3
// - Solve question 4: Earn 5 points
// Total points earned: 2 + 5 = 7. There is no other way to earn 7 or more points.

func MostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = int64(questions[i][0])
		if i+questions[i][1]+1 < n && i != n-1 {
			dp[i] += dp[i+questions[i][1]+1]
		}
		if dp[i] < dp[i+1] {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}
