package easy

import "sort"

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

// Use inorder traversal + sort:
// Time Complexity: O(n+nlogn)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 18.18%
// Memory Usage: 2 MB, less than 63.64%
func findSecondMinimumValue2(root *TreeNode) int {
	arr := inorder(root)
	sort.Ints(arr)

	var inorder func(root *TreeNode) []int
	inorder = func(root *TreeNode) []int {
		result := []int{}
		if root == nil {
			return result
		}
		result = append(result, inorder(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorder(root.Right)...)

		return result
	}

	min := arr[0]
	for _, v := range arr {
		if v > min {
			return v
		}
	}

	return -1
}
