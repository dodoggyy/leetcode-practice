package medium

// Use BFS:
// M: length of edges
// Q: length of queries
// L: avg. string length
// N: points
// Time Complexity: O(ML+Q⋅(L+M))
// Space Complexity:O(NL+M)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.08 MB, less than 89.90%
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	type Node struct {
		Name  int
		Ratio float64
	}

	result := make([]float64, len(queries))

	visit := map[string]int{}
	for _, v := range equations {
		if _, ok := visit[v[0]]; !ok {
			visit[v[0]] = len(visit)
		}
		if _, ok := visit[v[1]]; !ok {
			visit[v[1]] = len(visit)
		}
	}

	graph := make([][]Node, len(visit))
	for i, v := range equations {
		graph[visit[v[0]]] = append(graph[visit[v[0]]], Node{Name: visit[v[1]], Ratio: values[i]})
		graph[visit[v[1]]] = append(graph[visit[v[1]]], Node{Name: visit[v[0]], Ratio: 1.0 / values[i]})
	}

	bfs := func(start, end int) float64 {
		ratios := make([]float64, len(graph))
		ratios[start] = 1
		queue := []int{start}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if v == end {
				return ratios[v]
			}
			for _, e := range graph[v] {
				w := e.Name
				if ratios[w] == 0 {
					ratios[w] = ratios[v] * e.Ratio
					queue = append(queue, w)
				}
			}
		}

		return -1
	}

	for i, v := range queries {
		start, ok1 := visit[v[0]]
		end, ok2 := visit[v[1]]
		if !ok1 || !ok2 {
			result[i] = -1
		} else {
			result[i] = bfs(start, end)
		}

	}

	return result
}
