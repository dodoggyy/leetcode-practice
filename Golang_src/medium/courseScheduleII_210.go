package medium

// Use Topological Sort:
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
