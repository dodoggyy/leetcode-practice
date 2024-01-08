package easy

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 81 ms, faster than 48.28%
// Memory Usage: 7.08 MB, less than 100.00%
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
