package easy

// Use dfs:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 115 ms, faster than 47.78%
// Memory Usage: 14.60 MB, less than 59.44%
func longestZigZag(root *TreeNode) int {
	result := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(*TreeNode, int, bool)
	dfs = func(node *TreeNode, cur int, isLeft bool) {
		result = max(result, cur)
		if node == nil {
			return
		}
		if isLeft {
			if node.Left != nil {
				dfs(node.Left, 1, true)
			}
			if node.Right != nil {
				dfs(node.Right, cur+1, false)
			}
		} else {
			if node.Left != nil {
				dfs(node.Left, cur+1, true)
			}
			if node.Right != nil {
				dfs(node.Right, 1, false)
			}
		}
	}

	dfs(root, 0, true)
	dfs(root, 0, false)

	return result
}
