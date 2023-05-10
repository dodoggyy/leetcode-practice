package medium

// Use DFS:
// Time Complexity: O(E+V)
// Space Complexity:O(V)
// Runtime: 17 ms, faster than 90.41%
// Memory Usage: 6.9 MB, less than 52.5%
func isBipartite(graph [][]int) bool {
	// 0: unknown, 1: red, -1: blue
	colors := make([]int, len(graph))

	var dfs func(curNode, color int) bool
	dfs = func(curNode, color int) bool {
		if colors[curNode] != 0 {
			return colors[curNode] == color
		}
		colors[curNode] = color
		for _, v := range graph[curNode] {
			if !dfs(v, -color) {
				return false
			}
		}
		return true
	}

	for i := range graph {
		if colors[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}

// Use BFS:
// Time Complexity: O(E+V)
// Space Complexity:O(E+V)
// Runtime: 24 ms, faster than 86.08%
// Memory Usage: 6.1 MB, less than 27.85%
func isBipartite2(graph [][]int) bool {
	// 0: unknown, 1: red, -1: blue
	colors := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if colors[i] != 0 {
			continue
		}
		queue := []int{i}
		colors[i] = 1
		for len(queue) > 0 {
			curNode := queue[0]
			queue = queue[1:]
			for _, nextNode := range graph[curNode] {
				if colors[nextNode] == 0 {
					colors[nextNode] = -colors[curNode]
					queue = append(queue, nextNode)
				}
				if colors[nextNode] == colors[curNode] {
					return false
				}
			}
		}
	}

	return true
}
