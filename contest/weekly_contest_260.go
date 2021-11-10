package contest

/// Maximum Difference Between Increasing Elements
// Given a 0-indexed integer array nums of size n,
// find the maximum difference between nums[i] and nums[j] (i.e., nums[j] - nums[i]),
// such that 0 <= i < j < n and nums[i] < nums[j].
// Return the maximum difference. If no such i and j exists, return -1.
// Input: nums = [7,1,5,4]
// Output: 4
// Explanation:
// The maximum difference occurs with i = 1 and j = 2, nums[j] - nums[i] = 5 - 1 = 4.
// Note that with i = 1 and j = 0, the difference nums[j] - nums[i] = 7 - 1 = 6, but i > j, so it is not valid.
// Input: nums = [9,4,3,2]
// Output: -1
// Explanation:
// There is no i and j such that i < j and nums[i] < nums[j].
// Input: nums = [1,5,2,10]
// Output: 9
// Explanation:
// The maximum difference occurs with i = 0 and j = 3, nums[j] - nums[i] = 10 - 1 = 9.

func MaximumDifference(nums []int) int {
	result := 0
	return result
}

/// Grid Game
// You are given a 0-indexed 2D array grid of size 2 x n, where grid[r][c] represents the number of points at position (r, c) on the matrix. Two robots are playing a game on this matrix.
// Both robots initially start at (0, 0) and want to reach (1, n-1). Each robot may only move to the right ((r, c) to (r, c + 1)) or down ((r, c) to (r + 1, c)).
// At the start of the game, the first robot moves from (0, 0) to (1, n-1), collecting all the points from the cells on its path. For all cells (r, c) traversed on the path, grid[r][c] is set to 0.
// Then, the second robot moves from (0, 0) to (1, n-1),
// collecting the points on its path. Note that their paths may intersect with one another.
// The first robot wants to minimize the number of points collected by the second robot. In contrast, the second robot wants to maximize the number of points it collects. If both robots play optimally, return the number of points collected by the second robot.
// Input: grid = [[2,5,4],[1,5,1]]
// Output: 4
// Explanation: The optimal path taken by the first robot is shown in red,
// and the optimal path taken by the second robot is shown in blue.
// The cells visited by the first robot are set to 0.
// The second robot will collect 0 + 0 + 4 + 0 = 4 points.
// Input: grid = [[3,3,1],[8,5,2]]
// Output: 4
// Explanation: The optimal path taken by the first robot is shown in red,
// and the optimal path taken by the second robot is shown in blue.
// The cells visited by the first robot are set to 0.
// The second robot will collect 0 + 3 + 1 + 0 = 4 points.
// Input: grid = [[1,3,1,15],[1,3,3,1]]
// Output: 7
// Explanation: The optimal path taken by the first robot is shown in red,
// and the optimal path taken by the second robot is shown in blue.
// The cells visited by the first robot are set to 0.
// The second robot will collect 0 + 1 + 3 + 3 + 0 = 7 points.

func GridGame(grid [][]int) int64 {
	var result int64 = 0
	return result
}

/// Check if Word Can Be Placed In Crossword

// You are given an m x n matrix board, representing the current state of a crossword puzzle.
// The crossword contains lowercase English letters (from solved words),
// ' ' to represent any empty cells, and '#' to represent any blocked cells.
// A word can be placed horizontally (left to right or right to left) or vertically (top to bottom or bottom to top) in the board if:
// - It does not occupy a cell containing the character '#'.
// - The cell each letter is placed in must either be ' ' (empty) or match the letter already on the board.
// - There must not be any empty cells ' ' or other lowercase letters directly left or right of the word if the word was placed horizontally.
// - There must not be any empty cells ' ' or other lowercase letters directly above or below the word if the word was placed vertically.
// Given a string word, return true if word can be placed in board, or false otherwise.
// Input: board = [["#", " ", "#"], [" ", " ", "#"], ["#", "c", " "]], word = "abc"
// Output: true
// Explanation: The word "abc" can be placed as shown above (top to bottom).
// Input: board = [[" ", "#", "a"], [" ", "#", "c"], [" ", "#", "a"]], word = "ac"
// Output: false
// Explanation: It is impossible to place the word because there will always be a space/letter above or below it.
// Input: board = [["#", " ", "#"], [" ", " ", "#"], ["#", " ", "c"]], word = "ca"
// Output: true
// Explanation: The word "ca" can be placed as shown above (right to left).

func PlaceWordInCrossword(board [][]byte, word string) bool {
	return true
}

/// The Score of Students Solving Math Expression
// You are given a string s that contains digits 0-9, addition symbols '+',
// and multiplication symbols '*' only, representing a valid math expression of single digit numbers (e.g., 3+5*2).
// This expression was given to n elementary school students.
// The students were instructed to get the answer of the expression by following this order of operations:
// 1. Compute multiplication, reading from left to right; Then,
// 2. Compute addition, reading from left to right.
// You are given an integer array answers of length n,which are the submitted answers of the students in no particular order.
// You are asked to grade the answers, by following these rules:
// - If an answer equals the correct answer of the expression, this student will be rewarded 5 points;
// - Otherwise, if the answer could be interpreted as if the student applied the operators in the wrong order but had correct arithmetic,
// this student will be rewarded 2 points;
// - Otherwise, this student will be rewarded 0 points.
// Return the sum of the points of the students.
// Input: s = "7+3*1*2", answers = [20,13,42]
// Output: 7
// Explanation: As illustrated above, the correct answer of the expression is 13,
// therefore one student is rewarded 5 points: [20,13,42]
// A student might have applied the operators in this wrong order: ((7+3)*1)*2 = 20.
// Therefore one student is rewarded 2 points: [20,13,42]
// The points for the students are: [2,5,0]. The sum of the points is 2+5+0=7.
// Input: s = "3+5*2", answers = [13,0,10,13,13,16,16]
// Output: 19
// Explanation: The correct answer of the expression is 13,
// therefore three students are rewarded 5 points each: [13,0,10,13,13,16,16]
// A student might have applied the operators in this wrong order: ((3+5)*2 = 16.
// Therefore two students are rewarded 2 points: [13,0,10,13,13,16,16]
// The points for the students are: [5,0,0,5,5,2,2]. The sum of the points is 5+0+0+5+5+2+2=19.
// Input: s = "6+0*1", answers = [12,9,6,4,8,6]
// Output: 10
// Explanation: The correct answer of the expression is 6.
// If a student had incorrectly done (6+0)*1, the answer would also be 6.
// By the rules of grading, the students will still be rewarded 5 points (as they got the correct answer), not 2 points.
// The points for the students are: [0,0,5,0,0,5]. The sum of the points is 10.
// Input: s = "1+2*3+4", answers = [13,21,11,15]
// Output: 11
// Explanation: The correct answer of the expression is 11.
// Every other student was rewarded 2 points because they could have applied the operators as follows:
// - ((1+2)*3)+4 = 13
// - (1+2)*(3+4) = 21
// - 1+(2*(3+4)) = 15
// The points for the students are: [2,2,5,2]. The sum of the points is 11.

func ScoreOfStudents(s string, answers []int) int {
	return 0
}
