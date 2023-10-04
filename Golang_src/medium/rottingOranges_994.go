package medium

import "math"

// Use BFS 2:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 4 ms, faster than 62.40%
// Memory Usage: 3.0 MB, less than 22.56%
func orangesRotting2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	step := 0
	queue := [][]int{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				cnt++
				if grid[i][j] == 2 {
					queue = append(queue, []int{i, j})
				}
			}
		}
	}
	for len(queue) > 0 {
		size := len(queue)
		cnt -= size
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node[0] > 0 && grid[node[0]-1][node[1]] == 1 {
				grid[node[0]-1][node[1]] = 2
				queue = append(queue, []int{node[0] - 1, node[1]})
			}
			if node[0] < m-1 && grid[node[0]+1][node[1]] == 1 {
				grid[node[0]+1][node[1]] = 2
				queue = append(queue, []int{node[0] + 1, node[1]})
			}
			if node[1] > 0 && grid[node[0]][node[1]-1] == 1 {
				grid[node[0]][node[1]-1] = 2
				queue = append(queue, []int{node[0], node[1] - 1})
			}
			if node[1] < n-1 && grid[node[0]][node[1]+1] == 1 {
				grid[node[0]][node[1]+1] = 2
				queue = append(queue, []int{node[0], node[1] + 1})
			}
		}
		if len(queue) > 0 {
			step++
		}
	}

	if cnt > 0 {
		return -1
	}
	return step
}

// Use BFS:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 82.44%
// Memory Usage: 2.8 MB, less than 78.78%
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	hasFresh := false

	type Dim struct {
		X, Y int
	}

	queue := []Dim{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, Dim{i, j})
			}
			if !hasFresh && grid[i][j] == 1 {
				hasFresh = true
			}
		}
	}

	if !hasFresh {
		return 0
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		x, y := node.X, node.Y
		cur := grid[x][y]
		if x-1 >= 0 && grid[x-1][y] == 1 {
			grid[x-1][y] = cur + 1
			queue = append(queue, Dim{x - 1, y})
		}
		if x+1 < m && grid[x+1][y] == 1 {
			grid[x+1][y] = cur + 1
			queue = append(queue, Dim{x + 1, y})
		}
		if y-1 >= 0 && grid[x][y-1] == 1 {
			grid[x][y-1] = cur + 1
			queue = append(queue, Dim{x, y - 1})
		}
		if y+1 < n && grid[x][y+1] == 1 {
			grid[x][y+1] = cur + 1
			queue = append(queue, Dim{x, y + 1})
		}
	}

	result := math.MinInt32
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
			result = max(result, grid[i][j])
		}
	}

	return result - 2
}
