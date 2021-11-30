package top_question

/// Maximal Rectangle
// Given a rows x cols binary matrix filled with 0's and 1's,
// find the largest rectangle containing only 1's and return its area.
// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// Output: 6
// Explanation: The maximal rectangle is shown in the above picture.
// Input: matrix = []
// Output: 0
// Input: matrix = [["0"]]
// Output: 0
// Input: matrix = [["1"]]
// Output: 1
// Input: matrix = [["0","0"]]
// Output: 0

func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxArea, ans := 0, make([]int, len(matrix[0]))
	for _, row := range matrix {
		for j, value := range row {
			if value == '1' {
				ans[j]++
			} else {
				ans[j] = 0
			}
		}
		maxAreaInHistogram := findMaxAreaInHistogram(ans)
		if maxArea < maxAreaInHistogram {
			maxArea = maxAreaInHistogram
		}
	}
	return maxArea
}

func findMaxAreaInHistogram(ans []int) int {
	n, total := len(ans), 0
	stack := []int{-1}
	for i := 0; i <= n; i++ {
		var val int
		if i == n {
			val = -1
		} else {
			val = ans[i]
		}
		for stack[len(stack)-1] >= 0 && ans[stack[len(stack)-1]] > val {
			height := ans[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			area := height * width
			if total < area {
				total = area
			}
		}
		stack = append(stack, i)
	}
	return total
}
