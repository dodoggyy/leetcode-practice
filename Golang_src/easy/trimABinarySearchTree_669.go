package easy

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 55.00%
// Memory Usage: 6.2 MB, less than 5.00%
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
