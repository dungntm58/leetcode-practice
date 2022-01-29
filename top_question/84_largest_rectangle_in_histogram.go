package top_question

/// Largest Rectangle in Histogram
// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1,
// return the area of the largest rectangle in the histogram.
// Input: heights = [2,1,5,6,2,3]
// Output: 10
// Explanation: The above is a histogram where width of each bar is 1.
// The largest rectangle is shown in the red area, which has an area = 10 units.
// Input: heights = [2,4]
// Output: 4

func LargestRectangleArea(heights []int) int {
	max, n := 0, len(heights)
	stack := []int{-1}
	for i := 0; i <= n; i++ {
		val := 0
		if i == n {
			val = -1
		} else {
			val = heights[i]
		}
		for stack[len(stack)-1] >= 0 && heights[stack[len(stack)-1]] > val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := heights[top] * (i - stack[len(stack)-1] - 1)
			if area > max {
				max = area
			}
		}
		stack = append(stack, i)
	}
	return max
}
