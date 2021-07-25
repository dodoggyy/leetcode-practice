package medium

// Use DFS:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 288 ms, faster than 98.82%
// Memory Usage: 15.7 MB, less than 69.41%
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	result := 0

	var dfs func(row, col int)
	dfs = func(row, col int) {
		// ignore grid1 index = 1, grid2 index = 0 case
		if grid2[row][col] == 0 {
			return
		}
		if grid2[row][col] == 1 {
			grid2[row][col] = 0
		}
		if row-1 >= 0 && grid2[row-1][col] == 1 {
			dfs(row-1, col)
		}
		if row+1 < m && grid2[row+1][col] == 1 {
			dfs(row+1, col)
		}
		if col-1 >= 0 && grid2[row][col-1] == 1 {
			dfs(row, col-1)
		}
		if col+1 < n && grid2[row][col+1] == 1 {
			dfs(row, col+1)
		}
	}

	// floodfill to iterate over the islands of the second grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] != grid2[i][j] {
				dfs(i, j)
			}
		}
	}

	// check island in grid2
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				result++
				dfs(i, j)
			}
		}
	}

	return result
}
