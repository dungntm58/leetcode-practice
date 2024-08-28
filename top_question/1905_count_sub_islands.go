package top_question

// You are given two m x n binary matrices grid1 and grid2 containing only 0's (representing water) and 1's (representing land). An island is a group of 1's connected 4-directionally (horizontal or vertical). Any cells outside of the grid are considered water cells.
// An island in grid2 is considered a sub-island if there is an island in grid1 that contains all the cells that make up this island in grid2.
// Return the number of islands in grid2 that are considered sub-islands.

// Example 1:

// Input: grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
// Output: 3
// Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
// The 1s colored red in grid2 are those considered to be part of a sub-island. There are three sub-islands.

// Example 2:

// Input: grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
// Output: 2
// Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
// The 1s colored red in grid2 are those considered to be part of a sub-island. There are two sub-islands.

// Constraints:

// - m == grid1.length == grid2.length
// - n == grid1[i].length == grid2[i].length
// - 1 <= m, n <= 500
// - grid1[i][j] and grid2[i][j] are either 0 or 1.

func CountSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 0 {
				grid1[i][j] = 0
			} else {
				grid1[i][j] += grid2[i][j]
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 2 && bfsSubIslands(grid1, i, j, m, n) {
				res++
			}
		}
	}
	return res
}

func bfsSubIslands(grid [][]int, i, j, m, n int) bool {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return true
	}
	val := grid[i][j]
	if val > 0 {
		grid[i][j] = 0
	}
	r1 := bfsSubIslands(grid, i+1, j, m, n)
	r2 := bfsSubIslands(grid, i-1, j, m, n)
	r3 := bfsSubIslands(grid, i, j+1, m, n)
	r4 := bfsSubIslands(grid, i, j-1, m, n)
	return val == 2 && r1 && r2 && r3 && r4
}
