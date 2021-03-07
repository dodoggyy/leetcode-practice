package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 20 ms, faster than 76.51%
// Memory Usage: 9.7 MB, less than 53.02%
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	idx := findIndex(inorder, root.Val)

	// in   : L D R
	// post : L R D
	root.Left = buildTree(inorder[:idx], postorder[:len(inorder[:idx])])
	root.Right = buildTree(inorder[idx+1:], postorder[len(inorder[:idx]):len(postorder)-1])

	return root
}

func findIndex(order []int, target int) int {
	for i := range order {
		if order[i] == target {
			return i
		}
	}

	return -1
}
