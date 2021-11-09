package easy

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 98.22%
// Memory Usage: 5.7 MB, less than 100.00%
func isBalanced(root *TreeNode) bool {
	_, ok := depth2(root)

	return ok
}

func depth2(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, ok1 := depth2(root.Left)
	r, ok2 := depth2(root.Right)

	if !ok1 || !ok2 {
		return 0, false
	}

	if abs(l-r) > 1 {
		return 0, false
	}

	return 1 + max(l, r), true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
