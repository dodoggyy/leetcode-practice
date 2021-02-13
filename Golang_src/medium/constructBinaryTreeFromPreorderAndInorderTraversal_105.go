package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 36 ms, faster than 23.67%
// Memory Usage: 24.3 MB, less than 13.78%
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])

	return root
}
