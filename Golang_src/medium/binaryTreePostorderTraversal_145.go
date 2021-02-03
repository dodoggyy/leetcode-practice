package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 36.18%
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	postOrder := []int{}
	postOrder = append(postOrder, postorderTraversal(root.Left)...)
	postOrder = append(postOrder, postorderTraversal(root.Right)...)
	postOrder = append(postOrder, root.Val)

	return postOrder
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 36.18%
func postorderTraversal2(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}
		result = append(result, node.Val)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	size := len(result)
	for i := 0; i < size/2; i++ {
		result[i], result[size-1-i] = result[size-1-i], result[i]
	}

	return result
}
