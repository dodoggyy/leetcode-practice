package medium

// Use DFS:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 24 ms, faster than 20.33%
// Memory Usage: 5.2 MB, less than 66.43%
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	result := 0

	tmp := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		grid[r][c] = '0'
		tmp++

		if r-1 >= 0 && grid[r-1][c] == 1 {
			dfs(r-1, c)
		}
		if r+1 < m && grid[r+1][c] == 1 {
			dfs(r+1, c)
		}
		if c-1 >= 0 && grid[r][c-1] == 1 {
			dfs(r, c-1)
		}
		if c+1 < n && grid[r][c+1] == 1 {
			dfs(r, c+1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				if result < tmp {
					result = tmp
				}
				tmp = 0
			}
		}
	}

	return result
}
