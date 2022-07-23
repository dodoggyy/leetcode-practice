package medium

// Use BFS:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 88 ms, faster than 58.57%
// Memory Usage: 8 MB, less than 61.19%
func updateMatrix(mat [][]int) [][]int {
	defaultVal := -1

	m, n := len(mat), len(mat[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := range result {
		for j := range result[i] {
			if mat[i][j] != 0 {
				result[i][j] = defaultVal
			}
		}
	}

	type Idx struct {
		X, Y int
	}
	queue := []Idx{}

	for i := range result {
		for j := range result[i] {
			if result[i][j] == 0 {
				queue = append(queue, Idx{i, j})
			}
		}
	}

	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		cur := result[tmp.X][tmp.Y] + 1

		if tmp.X-1 >= 0 && result[tmp.X-1][tmp.Y] == defaultVal {
			result[tmp.X-1][tmp.Y] = cur
			queue = append(queue, Idx{tmp.X - 1, tmp.Y})
		}
		if tmp.X+1 < m && result[tmp.X+1][tmp.Y] == defaultVal {
			result[tmp.X+1][tmp.Y] = cur
			queue = append(queue, Idx{tmp.X + 1, tmp.Y})
		}
		if tmp.Y-1 >= 0 && result[tmp.X][tmp.Y-1] == defaultVal {
			result[tmp.X][tmp.Y-1] = cur
			queue = append(queue, Idx{tmp.X, tmp.Y - 1})
		}
		if tmp.Y+1 < n && result[tmp.X][tmp.Y+1] == defaultVal {
			result[tmp.X][tmp.Y+1] = cur
			queue = append(queue, Idx{tmp.X, tmp.Y + 1})
		}

	}

	return result
}
