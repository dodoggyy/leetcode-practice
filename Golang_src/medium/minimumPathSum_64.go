package medium

// Use DP:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 85.32%
// Memory Usage: 3.9 MB, less than 79.86%
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i > 0 && j > 0 {
				grid[i][j] = min(grid[i][j]+grid[i-1][j], grid[i][j]+grid[i][j-1])
			} else if i > 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else if j > 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}
		}
	}

	return grid[m-1][n-1]
}
