package medium

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

// Use stack:
// Time Complexity: O(h+k)
// Space Complexity:O(h+k)
// h: tree height, k: k smallest element
// Runtime: 8 ms, faster than 97.33%
// Memory Usage: 6 MB, less than 38.17%
func kthSmallest2(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	cur := root

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}

	return 0
}
