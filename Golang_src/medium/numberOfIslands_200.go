package medium

// Use DFS:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.9 MB, less than 68.49%
func numIslands(grid [][]byte) int {
	result := 0
	m, n := len(grid), len(grid[0])

	var dfs func(m, n int)
	dfs = func(row, col int) {
		grid[row][col] = '0'
		if row-1 >= 0 && grid[row-1][col] == '1' {
			dfs(row-1, col)
		}
		if row+1 < m && grid[row+1][col] == '1' {
			dfs(row+1, col)
		}
		if col-1 >= 0 && grid[row][col-1] == '1' {
			dfs(row, col-1)
		}
		if col+1 < n && grid[row][col+1] == '1' {
			dfs(row, col+1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				dfs(i, j)
			}
		}
	}

	return result
}

// Use BFS:
// Time Complexity: O(m*n)
// Space Complexity: O(min(m,n))
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.6 MB, less than 33.98%
func numIslands2(grid [][]byte) int {
	result := 0
	m, n := len(grid), len(grid[0])

	type point struct {
		x int
		y int
	}

	queue := []*point{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				grid[i][j] = '0'
				queue = append(queue, &point{x: i, y: j})
				for len(queue) > 0 {
					tmp := queue[0]
					queue = queue[1:]
					row, col := tmp.x, tmp.y
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, &point{x: row - 1, y: col})
						grid[row-1][col] = '0'
					}
					if row+1 < m && grid[row+1][col] == '1' {
						queue = append(queue, &point{x: row + 1, y: col})
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, &point{x: row, y: col - 1})
						grid[row][col-1] = '0'
					}
					if col+1 < n && grid[row][col+1] == '1' {
						queue = append(queue, &point{x: row, y: col + 1})
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}

	return result
}
