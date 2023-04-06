package medium

// Use DFS:
// Time Complexity: O(m*n)
// Space Complexity:O(m+n)
// Runtime: 8 ms, faster than 80.85%
// Memory Usage: 4.8 MB, less than 55.32%
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	result := 0

	var dfs func(i, j int, result bool) bool
	dfs = func(i, j int, result bool) bool {
		if i == 0 || j == 0 || i == m-1 || j == n-1 {
			return false
		}

		grid[i][j] = 2

		if i > 0 && grid[i-1][j] == 0 {
			result = dfs(i-1, j, result) && result
		}
		if i < m-1 && grid[i+1][j] == 0 {
			result = dfs(i+1, j, result) && result
		}
		if j > 0 && grid[i][j-1] == 0 {
			result = dfs(i, j-1, result) && result
		}
		if j < n-1 && grid[i][j+1] == 0 {
			result = dfs(i, j+1, result) && result
		}

		return result
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 0 {
				if dfs(i, j, true) {
					result++
				}
			}
		}
	}

	return result
}
