package top_question

/// Spiral Matrix
// Given an m x n matrix, return all elements of the matrix in spiral order.
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)
	l, r, u, d := 0, 0, 0, 0
	direction := 1
	index := 0
	row, col := 0, 0
	if n == 1 {
		direction = 2
	}
	for index < m*n {
		ans[index] = matrix[row][col]
		switch direction {
		case 1:
			col++
			if col == n-r-1 {
				u++
				direction++
			}
		case 2:
			row++
			if row == m-d-1 {
				r++
				direction++
			}
		case 3:
			col--
			if col == l {
				d++
				direction++
			}
		case 4:
			row--
			if row == u {
				l++
				direction = 1
			}
		}
		index++
	}
	return ans
}
