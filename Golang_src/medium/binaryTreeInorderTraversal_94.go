package medium

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 22.04%
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	inorder := []int{}
	inorder = append(inorder, inorderTraversal(root.Left)...)
	inorder = append(inorder, root.Val)
	inorder = append(inorder, inorderTraversal(root.Right)...)

	return inorder
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 22.04%
func inorderTraversal2(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}

	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}

	return result
}
