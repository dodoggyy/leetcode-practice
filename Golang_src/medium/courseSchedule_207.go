package medium

// Use Topological Sort:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 96.25%
// Memory Usage: 6.1 MB, less than 41.57%
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	// state: 0 = unknown, 1 = visiting, 2 = visited
	visit := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if hasCycle(i, &visit, &graph) {
			return false
		}
	}

	return true
}

func hasCycle(cur int, visit *[]int, graph *[][]int) bool {
	if (*visit)[cur] == 1 { // 1 = visiting
		return true
	}
	if (*visit)[cur] == 2 { // 2 = visited
		return false
	}

	(*visit)[cur] = 1 // 1 = visiting

	for next := 0; next < len((*graph)[cur]); next++ {
		if hasCycle((*graph)[cur][next], visit, graph) {
			return true
		}
	}

	(*visit)[cur] = 2 // 2 = visited

	return false
}
