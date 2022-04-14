package easy

// Use recursive:
// Time Complexity: O(H)
// Space Complexity:O(1)
// H for tree height
// Runtime: 16 ms, faster than 99.40%
// Memory Usage: 7 MB, less than 95.77%
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return searchBST(root.Left, val)
}
