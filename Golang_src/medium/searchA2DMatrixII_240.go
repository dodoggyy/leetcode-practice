package medium

// Use Search Space Reduction:
// m: row
// n: col
// Time Complexity: O(m+n)
// Space Complexity:O(l)
// Runtime: 34 ms, faster than 49.80%
// Memory Usage: 7.6 MB, less than 13.04%
func searchMatrixs(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])

	i, j := 0, col-1
	for i < row && j >= 0 {
		cur := matrix[i][j]

		if cur > target {
			j--
		} else if cur < target {
			i++
		} else {
			return true
		}
	}

	return false
}
