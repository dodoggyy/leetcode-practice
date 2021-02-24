package medium

// Node Definition
type Node struct {
	Val       int
	Neighbors []*Node
}

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.7 MB, less than 27.23%
func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}

		// if node has visited, then return exist clone node from hash table
		if _, exist := visited[node]; exist {
			return visited[node]
		}

		cloneNode := &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
		// store to hash table
		visited[node] = cloneNode

		// go through all neighbors' node and update clone node neighbors
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}

		return cloneNode
	}

	return cg(node)
}

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.7 MB, less than 27.23%
func cloneGraph2(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := map[*Node]*Node{}

	queue := []*Node{node}
	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		for _, neighbor := range tmp.Neighbors {
			if _, exist := visited[neighbor]; !exist {
				visited[neighbor] = &Node{
					Val:       neighbor.Val,
					Neighbors: []*Node{},
				}
				queue = append(queue, neighbor)
			}
			visited[tmp].Neighbors = append(visited[tmp].Neighbors, visited[neighbor])
		}

	}

	return visited[node]
}
