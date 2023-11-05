package easy

import "math"

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

// Use dfs compare current one with previous one to reduce space:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 10 ms, faster than 67.52%
// Memory Usage: 6.54 MB, less than 79.30%
func getMinimumDifference2(root *TreeNode) int {
	result := math.MaxInt32
	pre := -1

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			if node.Val-pre < result {
				result = node.Val - pre
			}
		}
		pre = node.Val
		dfs(node.Right)
	}

	dfs(root)

	return result
}
