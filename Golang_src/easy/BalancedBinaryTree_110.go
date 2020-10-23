package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 98.22%
// Memory Usage: 5.7 MB, less than 100.00%
func isBalanced(root *TreeNode) bool {
	_, ok := depth(root)

	return ok
}

func depth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, ok1 := depth(root.Left)
	r, ok2 := depth(root.Right)

	if !ok1 || !ok2 {
		return 0, false
	}

	if abs(l-r) > 1 {
		return 0, false
	}

	return 1 + max(l, r), true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
