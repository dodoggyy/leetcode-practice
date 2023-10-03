package medium

// Use BFS:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 182 ms, faster than 72.67%
// Memory Usage: 14.80 MB, less than 90.70%
func minReorder(n int, connections [][]int) int {
	result := 0

	queue := []int{0}
	visit := make([]bool, n)
	visit[0] = true

	adjList := map[int][]int{}
	for i := range connections {
		adjList[connections[i][0]] = append(adjList[connections[i][0]], connections[i][1])
		adjList[connections[i][1]] = append(adjList[connections[i][1]], -connections[i][0])
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, v := range adjList[cur] {
				if !visit[abs(v)] {
					if v > 0 {
						result++
					}
					v = abs(v)
					visit[v] = true
					queue = append(queue, v)

				}
			}
		}
	}

	return result
}
