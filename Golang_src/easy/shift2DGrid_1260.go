package easy

// Use Modulo Arithmetic:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 63 ms, faster than 5.26%
// Memory Usage: 8.4 MB, less than 7.89%
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[(i+((j+k)/n))%m][(j+k)%n] = grid[i][j]
		}
	}

	return result
}
