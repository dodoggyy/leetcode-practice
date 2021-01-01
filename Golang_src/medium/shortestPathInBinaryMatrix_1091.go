package medium

// Use BFS:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 40 ms, faster than 98.59%
// Memory Usage: 6.4 MB, less than 100.00%
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	row, col := len(grid), len(grid[0])
	if row == 1 && col == 1 {
		return 1
	}

	dx, dy := []int{-1, 0, 1, 1, 1, 0, -1, -1}, []int{-1, -1, -1, 0, 1, 1, 1, 0}
	queue := []int{0}
	grid[0][0] = 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur/col, cur%col

		for i := range dx {
			nx, ny := x+dx[i], y+dy[i]
			if nx >= 0 && nx < row && ny >= 0 && ny < col && grid[nx][ny] == 0 {
				queue = append(queue, nx*col+ny)
				grid[nx][ny] = grid[x][y] + 1
				if nx == row-1 && ny == col-1 {
					return grid[nx][ny]
				}
			}
		}
	}

	return -1
}
