package easy

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(logn)
// Runtime: 4 ms, faster than 92.48%
// Memory Usage: 4.4 MB, less than 15.52%
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
