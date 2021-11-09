package hard

import "math"

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 16 ms, faster than 89.02%
// Memory Usage: 6.7 MB, less than 99.19%
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := math.MinInt32
	helper(root, &result)

	return result
}

func helper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	l := max(0, helper(root.Left, result))
	r := max(0, helper(root.Right, result))
	sum := l + r + root.Val
	*result = max(*result, sum)

	return max(l, r) + root.Val
}
