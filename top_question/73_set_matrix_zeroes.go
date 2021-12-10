package top_question

/// Set Matrix Zeroes
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's, and return the matrix.
// You must do it in place.
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
// Follow up:
// - A straightforward solution using O(mn) space is probably a bad idea.
// - A simple improvement uses O(m + n) space, but still not the best solution.
// - Could you devise a constant space solution?

func SetZeroes(matrix [][]int) {
	isCol := false
	row := len(matrix)
	col := len(matrix[0])
	for _, row := range matrix {
		if row[0] == 0 {
			isCol = true
		}

		for j := 1; j < col; j++ {
			if row[j] == 0 {
				row[0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		firstRow := matrix[0]
		for i := range firstRow {
			firstRow[i] = 0
		}
	}

	if isCol {
		for _, row := range matrix {
			row[0] = 0
		}
	}
}
