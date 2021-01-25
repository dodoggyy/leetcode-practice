package medium

// Use DFS:
// Time Complexity: O(E+V)
// Space Complexity:O(E+V)
// Runtime: 20 ms, faster than 100.00%
// Memory Usage: 6.1 MB, less than 92.41%
func isBipartite(graph [][]int) bool {
	// 0: unknown, 1: red, -1: blue
	colors := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if colors[i] == 0 && !dfs(i, 1, &colors, &graph) {
			return false
		}
	}
	return true
}

func dfs(curNode, curColor int, colors *[]int, graph *[][]int) bool {
	if (*colors)[curNode] != 0 {
		return (*colors)[curNode] == curColor
	}
	(*colors)[curNode] = curColor
	for _, nextNode := range (*graph)[curNode] {
		if !dfs(nextNode, -curColor, colors, graph) {
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
