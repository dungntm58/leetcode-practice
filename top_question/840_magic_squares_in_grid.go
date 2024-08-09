package top_question

// A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.
// Given a row x col grid of integers, how many 3 x 3 contiguous magic square subgrids are there?
// Note: while a magic square can only contain numbers from 1 to 9, grid may contain numbers up to 15.

// Example 1:

// 4 3 8 4
// 9 5 1 9
// 2 7 6 2

// Input: grid = [[4,3,8,4],[9,5,1,9],[2,7,6,2]]
// Output: 1
// Explanation:
// The following subgrid is a 3 x 3 magic square:

// 4 3 8
// 9 5 1
// 2 7 6

// while this one is not:

// 3 8 4
// 5 1 9
// 7 6 2

// In total, there is only one magic square inside the given grid.
// Example 2:

// Input: grid = [[8]]
// Output: 0

// Constraints:
// - row == grid.length
// - col == grid[i].length
// - 1 <= row, col <= 10
// - 0 <= grid[i][j] <= 15

func NumMagicSquaresInside(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row < 3 || col < 3 {
		return 0
	}
	count := 0
	for i := 0; i < row-3; i++ {
		for j := 0; j < col-3; j++ {
			if isMagicSquare(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func isMagicSquare(grid [][]int, rStart int, cStart int) bool {
	var numberCheck = [9]bool{}
	for i := rStart; i < rStart+3; i++ {
		for j := cStart; j < cStart+3; j++ {
			if grid[i][j] == 0 || grid[i][j] > 9 {
				return false
			}
			if numberCheck[grid[i][j]-1] {
				return false
			}
			numberCheck[grid[i][j]-1] = true
		}
	}

	var row1 = grid[rStart][cStart] + grid[rStart][cStart+1] + grid[rStart][cStart+2]
	var row2 = grid[rStart+1][cStart] + grid[rStart+1][cStart+1] + grid[rStart+1][cStart+2]

	if row1 != row2 {
		return false
	}

	var row3 = grid[rStart+2][cStart] + grid[rStart+2][cStart+1] + grid[rStart+2][cStart+2]
	if row2 != row3 {
		return false
	}

	var col1 = grid[rStart][cStart] + grid[rStart+1][cStart] + grid[rStart+2][cStart]
	var col2 = grid[rStart][cStart+1] + grid[rStart+1][cStart+1] + grid[rStart+2][cStart+1]

	if col1 != col2 {
		return false
	}

	var col3 = grid[rStart][cStart+2] + grid[rStart+1][cStart+2] + grid[rStart+2][cStart+2]
	if col2 != col3 {
		return false
	}

	var d1 = grid[rStart][cStart] + grid[rStart+1][cStart+1] + grid[rStart+2][cStart+2]
	var d2 = grid[rStart+2][cStart] + grid[rStart+1][cStart+1] + grid[rStart][cStart+2]

	return d1 == d2
}
