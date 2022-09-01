package medium

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 115 ms, faster than 64.07%
// Memory Usage: 10.2 MB, less than 97.60%
func goodNodes(root *TreeNode) int {
	result := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, maxCur int) {
		if node == nil {
			return
		}

		if node.Val >= maxCur {
			result++
			maxCur = node.Val
		}
		dfs(node.Left, maxCur)
		dfs(node.Right, maxCur)
	}

	dfs(root, root.Val)

	return result
}
