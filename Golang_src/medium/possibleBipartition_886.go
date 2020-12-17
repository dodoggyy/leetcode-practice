package medium

// Use DFS:
// Time Complexity: O(V+E)
// Space Complexity:O(V+E)
// Runtime: 136 ms, faster than 21.43%
// Memory Usage: 7.1 MB, less than 97.62%
func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]int, N)

	// O(V), avoid creating N*N array
	for _, v := range dislikes {
		graph[v[0]-1] = append(graph[v[0]-1], v[1]-1)
		graph[v[1]-1] = append(graph[v[1]-1], v[0]-1)
	}
	// 0: unknown, 1: red, -1: blue
	colors := make([]int, N)

	for i := 0; i < N; i++ {
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
// Time Complexity: O(V+E)
// Space Complexity:O(V+E)
// Runtime: 180 ms, faster than 11.90%
// Memory Usage: 7.1 MB, less than 97.62%
func possibleBipartition2(N int, dislikes [][]int) bool {
	graph := make([][]int, N)

	// O(V)
	for _, v := range dislikes {
		graph[v[0]-1] = append(graph[v[0]-1], v[1]-1)
		graph[v[1]-1] = append(graph[v[1]-1], v[0]-1)
	}
	// 0: unknown, 1: red, -1: blue
	colors := make([]int, N)

	for i := 0; i < N; i++ {
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
