package medium

// Use array counter:
// Time Complexity: O(m*n)
// Space Complexity:O(m+n)
// Runtime: 165 ms, faster than 100.00%
// Memory Usage: 13.98 MB, less than 100.00%
func onesMinusZeros(grid [][]int) [][]int {
	cntRow := make([]int, len(grid))
	cntCol := make([]int, len(grid[0]))
	for i := range grid {
		for j, v := range grid[i] {
			if v == 0 {
				cntRow[i]--
				cntCol[j]--
			} else if v == 1 {
				cntRow[i]++
				cntCol[j]++
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = cntRow[i] + cntCol[j]
		}
	}

	return grid
}
