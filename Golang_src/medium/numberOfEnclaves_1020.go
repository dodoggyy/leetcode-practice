package medium

// Use DFS:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 48 ms, faster than 74.14%
// Memory Usage: 6.8 MB, less than 29.31%
func numEnclaves(grid [][]int) int {
	result := 0

	m, n := len(grid), len(grid[0])
	cnt := 0

	var dfs func(r, c int, noBound bool) bool
	dfs = func(r, c int, noBound bool) bool {
		grid[r][c] = 2

		if r == 0 || c == 0 || r == m-1 || c == n-1 {
			noBound = false
		}

		if r-1 > 0 && grid[r-1][c] == 1 {
			cnt++
			noBound = dfs(r-1, c, noBound) && noBound
		}
		if r+1 < m && grid[r+1][c] == 1 {
			cnt++
			noBound = dfs(r+1, c, noBound) && noBound
		}
		if c-1 > 0 && grid[r][c-1] == 1 {
			cnt++
			noBound = dfs(r, c-1, noBound) && noBound
		}
		if c+1 < n && grid[r][c+1] == 1 {
			cnt++
			noBound = dfs(r, c+1, noBound) && noBound
		}

		return noBound
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				cnt = 1
				if dfs(i, j, true) {
					result += cnt
				}
			}
		}
	}

	return result
}
