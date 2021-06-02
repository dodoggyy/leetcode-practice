package medium

// Use iterative in-place:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 0.0 MB, less than 100%
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			cnt := 0

			// prevent cnt + 1 when origin element = 1
			if board[row][col] == 1 {
				cnt--
			}

			for r := -1; r < 2; r++ {
				for c := -1; c < 2; c++ {
					rr := row + r
					cc := col + c
					if (rr < m && rr >= 0) && (cc < n && cc >= 0) && abs289(board[rr][cc]) == 1 {
						cnt++
					}
				}
			}
			// live -> dead : mark as -1
			if board[row][col] == 1 && (cnt < 2 || cnt > 3) {
				board[row][col] = -1
			}

			// dead -> live : mark as 2
			if board[row][col] == 0 && cnt == 3 {
				board[row][col] = 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

func abs289(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
