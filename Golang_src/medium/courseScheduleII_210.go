package medium

// Use Topological Sort vis DFS:
// Time Complexity: O(V+E)
// Space Complexity:O(V+E)
// Runtime: 4 ms, faster than 100.00%
// Memory Usage: 6.1 MB, less than 50.00%
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	// state: 0 = unknown, 1 = visiting, 2 = visited
	visit := make([]int, numCourses)
	result := []int{}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i, &graph, &visit, &result) {
			return []int{}
		}
	}

	// need reverse slice
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result
}

func hasCycle(cur int, graph *[][]int, visit *[]int, result *[]int) bool {
	if (*visit)[cur] == 1 {
		return true
	}
	if (*visit)[cur] == 2 {
		return false
	}
	(*visit)[cur] = 1

	for next := 0; next < len((*graph)[cur]); next++ {
		if hasCycle((*graph)[cur][next], graph, visit, result) {
			return true
		}
	}

	(*visit)[cur] = 2
	*result = append(*result, cur)

	return false
}

// Use Topological Sort via BFS:
// Time Complexity: O(V+E)
// Space Complexity:O(V+E)
// Runtime: 11 ms, faster than 73.00%
// Memory Usage: 6.28 MB, less than 90.38%
func findOrder2(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	result := []int{}
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		indeg[v[0]]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, v := range edges[node] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(result) != numCourses {
		return []int{}
	}

	return result
}
