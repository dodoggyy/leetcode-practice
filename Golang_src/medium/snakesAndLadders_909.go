package medium

// Use BFS:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 5 ms, faster than 96.70%
// Memory Usage: 4.78 MB, less than 90.11%
func snakesAndLadders(board [][]int) int {
	n := len(board)

	idx2rc := func(idx int) (int, int) {
		r, c := (idx-1)/n, (idx-1)%n
		if r%2 == 1 {
			c = n - 1 - c
		}
		r = n - 1 - r

		return r, c
	}

	visit := make([]bool, n*n)
	type pair struct {
		idx, step int
	}
	queue := []pair{{1, 0}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for i := 1; i <= 6; i++ {
			next := node.idx + i
			if next > n*n {
				break
			}
			r, c := idx2rc(next)
			if board[r][c] > 0 {
				next = board[r][c]
			}
			if next == n*n {
				return node.step + 1
			}
			if !visit[next] {
				visit[next] = true
				queue = append(queue, pair{next, node.step + 1})
			}
		}
	}
	return -1
}
