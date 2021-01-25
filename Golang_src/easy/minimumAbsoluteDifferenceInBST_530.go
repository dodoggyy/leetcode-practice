package easy

import "math"

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use inorder:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 70.00%
// Memory Usage: 6.5 MB, less than 8.00%
func getMinimumDifference(root *TreeNode) int {
	order := inorder(root)
	result := math.MaxInt16

	for i := 1; i < len(order); i++ {
		result = min(result, order[i]-order[i-1])
	}

	return result
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	result = append(result, inorder(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorder(root.Right)...)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Use dfs compare current one with previous one to reduce space:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 12 ms, faster than 70.00%
// Memory Usage: 6 MB, less than 38.00%
func getMinimumDifference2(root *TreeNode) int {
	pre := -1
	result := math.MaxInt16

	dfs(root, &pre, &result)

	return result
}

func dfs(root *TreeNode, pre, min *int) {
	if root != nil {
		dfs(root.Left, pre, min)
		if *pre != -1 {
			if root.Val-*pre < *min {
				*min = root.Val - *pre
			}
		}
		*pre = root.Val
		dfs(root.Right, pre, min)
	}
}
