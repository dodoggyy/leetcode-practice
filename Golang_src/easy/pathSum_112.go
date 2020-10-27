package easy

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 95.74%
// Memory Usage: 4.8 MB, less than 100.00%
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && (sum-root.Val) == 0 {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
