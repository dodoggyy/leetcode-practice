package easy

import "fmt"

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 22 ms, faster than 47.83%
// Memory Usage: 8.2 MB, less than 58.70%
func tree2str(root *TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return fmt.Sprintf("%v", root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%v(%v)", root.Val, tree2str(root.Left))
	}

	return fmt.Sprintf("%v(%v)(%v)", root.Val, tree2str(root.Left), tree2str(root.Right))
}
