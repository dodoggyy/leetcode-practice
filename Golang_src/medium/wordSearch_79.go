package medium

// Use DFS & backtracking:
// m: row length, n: column length
// L: word length
// Time Complexity: O(m*n*3^L)
// Space Complexity:O(m*n)
// Runtime: 4 ms, faster than 96.97%
// Memory Usage: 3.8 MB, less than 36.36%
func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	visit := make([][]bool, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]bool, col)
	}

	var dfs func(curRow, curCol, step int) bool
	dfs = func(curRow, curCol, step int) bool {
		if step == len(word) {
			return true
		}
		if curRow < 0 || curRow >= row || curCol < 0 || curCol >= col {
			return false
		}
		if visit[curRow][curCol] || board[curRow][curCol] != word[step] {
			return false
		}
		visit[curRow][curCol] = true
		find := dfs(curRow+1, curCol, step+1) || dfs(curRow-1, curCol, step+1) || dfs(curRow, curCol+1, step+1) || dfs(curRow, curCol-1, step+1)

		if find {
			return true
		}
		visit[curRow][curCol] = false
		return false

	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
