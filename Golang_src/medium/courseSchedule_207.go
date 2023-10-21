package medium

// Use Topological Sort via DFS:
// Time Complexity: O(n+m)
// Space Complexity:O(n+m)
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
		if hasCycle2(i, &visit, &graph) {
			return false
		}
	}

	return true
}

func hasCycle2(cur int, visit *[]int, graph *[][]int) bool {
	if (*visit)[cur] == 1 { // 1 = visiting
		return true
	}
	if (*visit)[cur] == 2 { // 2 = visited
		return false
	}

	(*visit)[cur] = 1 // 1 = visiting

	for next := 0; next < len((*graph)[cur]); next++ {
		if hasCycle2((*graph)[cur][next], visit, graph) {
			return true
		}
	}

	(*visit)[cur] = 2 // 2 = visited

	return false
}

// Use Topological Sort via DFS:
// Time Complexity: O(n+m)
// Space Complexity:O(n+m)
// Runtime: 26 ms, faster than 23.60%
// Memory Usage: 6.4 MB, less than 77.00%
func canFinish2(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := range prerequisites {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	// state: 0 = not visit, 1 = visiting, 2 = visited
	visit := make([]int, numCourses)

	var hasCycle func(cur int) bool
	hasCycle = func(cur int) bool {
		if visit[cur] == 1 {
			return true
		}
		if visit[cur] == 2 {
			return false
		}

		visit[cur] = 1

		for i := 0; i < len(graph[cur]); i++ {
			if hasCycle(graph[cur][i]) {
				return true
			}
		}

		visit[cur] = 2

		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}

	return true
}

// Use Topological Sort via BFS:
// Time Complexity: O(n+m)
// Space Complexity:O(n+m)
// Runtime: 12 ms, faster than 68.60%
// Memory Usage: 6.05 MB, less than 97.11%
func canFinish3(numCourses int, prerequisites [][]int) bool {
	// take indegree = 0 to queue
	// then minus other point indegree
	// in the same time if indegree = 0 then append to queue
	// finally check topologic path equal total points or not

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

	return len(result) == numCourses
}
