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
// Runtime: 4 ms, faster than 94.69%
// Memory Usage: 4.5 MB, less than 100.00%
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	depth(root, &result)

	return result
}

func depth(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left, result)
	r := depth(root.Right, result)

	*result = max((l + r), *result)

	return 1 + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
