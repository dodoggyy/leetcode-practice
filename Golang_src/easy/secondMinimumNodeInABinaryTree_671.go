package easy

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 20.00%
func findSecondMinimumValue(root *TreeNode) int {
	return dfs2(root, root.Val)
}

func dfs2(root *TreeNode, target int) int {
	if root == nil {
		return -1
	}

	if root.Val > target {
		return root.Val
	}

	left := dfs2(root.Left, target)
	right := dfs2(root.Right, target)

	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}

	return min(left, right)
}
