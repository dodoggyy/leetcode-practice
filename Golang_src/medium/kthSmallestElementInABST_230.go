package medium

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use inorder traversal with memorize:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 16 ms, faster than 8.20%
// Memory Usage: 6.9 MB, less than 5.33%
func kthSmallest(root *TreeNode, k int) int {
	order := inorder(root)

	return order[k-1]
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
