package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 20.59%
// Memory Usage: 2.4 MB, less than 50.00%
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
