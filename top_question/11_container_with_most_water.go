package top_question

/// Container With Most Water
// Given n non-negative integers a1, a2, ..., an ,
// where each represents a point at coordinate (i, ai).
// n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
// Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
// Notice that you may not slant the container.
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
// In this case, the max area of water (blue section) the container can contain is 49.
// Input: height = [1,1]
// Output: 1
// Input: height = [4,3,2,1,4]
// Output: 16
// Input: height = [1,2,1]
// Output: 2

func MaxArea(height []int) int {
	left, right, r := 0, len(height)-1, 0
	for left < right {
		boundary := height[left]
		if boundary > height[right] {
			boundary = height[right]
		}
		newResult := boundary * (right - left)
		if newResult > r {
			r = newResult
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return r
}
