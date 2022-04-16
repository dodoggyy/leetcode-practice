package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 79.71%
// Memory Usage: 6.7 MB, less than 78.26%
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		sum += node.Val
		node.Val = sum
		traversal(node.Right)
	}

	traversal(root)

	return root
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 208 ms, faster than 60.78%
// Memory Usage: 212.8 MB, less than 60.78%
func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	cur := root
	stack := []*TreeNode{}

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Val += sum
		sum = cur.Val
		cur = cur.Left
	}

	return root
}
