package medium

// Use one pass iterative:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.63 MB, less than 57.83%
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	cols := [9][9]int{}
	areas := [3][3][9]int{}
	for i, r := range board {
		for j, ch := range r {
			if ch == '.' {
				continue
			}
			idx := ch - '1'
			rows[i][idx]++
			cols[j][idx]++
			areas[i/3][j/3][idx]++
			if rows[i][idx] > 1 || cols[j][idx] > 1 || areas[i/3][j/3][idx] > 1 {
				return false
			}
		}
	}

	return true
}
