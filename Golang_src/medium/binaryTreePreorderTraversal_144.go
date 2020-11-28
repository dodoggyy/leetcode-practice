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
// Memory Usage: 2.1 MB, less than 31.58%
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	preorder := []int{}
	preorder = append(preorder, root.Val)
	preorder = append(preorder, preorderTraversal(root.Left)...)
	preorder = append(preorder, preorderTraversal(root.Right)...)

	return preorder
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 31.58%
func preorderTraversal2(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		result = append(result, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}

	return result
}
