package top_question

/// Number of Islands
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
// return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1
// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3

func NumIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfsVisitIslands(grid, i, j, visited)
				count++
			}
		}
	}
	return count
}

var nbrRow = [4]int{-1, 0, 0, 1}
var nbrCol = [4]int{0, -1, 1, 0}

func dfsVisitIslands(grid [][]byte, i, j int, visited [][]bool) {
	visited[i][j] = true

	for x := 0; x < 4; x++ {
		newI, newJ := i+nbrRow[x], j+nbrCol[x]
		if newI >= 0 && newJ >= 0 && newI < len(grid) && newJ < len(grid[0]) && (grid[newI][newJ] == '1' && !visited[newI][newJ]) {
			dfsVisitIslands(grid, newI, newJ, visited)
		}
	}
}
