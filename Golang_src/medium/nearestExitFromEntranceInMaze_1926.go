package medium

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 6.94 MB, less than 97.01%
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])

	type Coordinate struct {
		X, Y int
		Step int
	}

	start := Coordinate{X: entrance[0], Y: entrance[1], Step: 0}

	queue := []Coordinate{start}

	maze[start.X][start.Y] = '+'

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x, y, step := node.X, node.Y, node.Step
		if (x == m-1 || x == 0 || y == n-1 || y == 0) && step > 0 {
			return step
		}

		if x > 0 && maze[x-1][y] == '.' {
			queue = append(queue, Coordinate{X: x - 1, Y: y, Step: step + 1})
			maze[x-1][y] = '+'
		}
		if x < m-1 && maze[x+1][y] == '.' {
			queue = append(queue, Coordinate{X: x + 1, Y: y, Step: step + 1})
			maze[x+1][y] = '+'
		}
		if y > 0 && maze[x][y-1] == '.' {
			queue = append(queue, Coordinate{X: x, Y: y - 1, Step: step + 1})
			maze[x][y-1] = '+'
		}
		if y < n-1 && maze[x][y+1] == '.' {
			queue = append(queue, Coordinate{X: x, Y: y + 1, Step: step + 1})
			maze[x][y+1] = '+'
		}

	}

	return -1
}
