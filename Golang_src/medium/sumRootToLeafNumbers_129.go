package medium

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.2 MB, less than 43.64%
func sumNumbers(root *TreeNode) int {
	return helperSumNumbers(root, 0)
}

func helperSumNumbers(root *TreeNode, pre int) int {
	if root == nil {
		return 0
	}
	sum := pre*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return helperSumNumbers(root.Left, sum) + helperSumNumbers(root.Right, sum)
}
