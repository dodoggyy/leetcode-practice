package easy

type Node struct {
	Val      int
	Children []*Node
}

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 61.10%
// Memory Usage: 4 MB, less than 65.54%
func preorderT(root *Node) []int {
	result := []int{}

	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}

	dfs(root)

	return result
}
