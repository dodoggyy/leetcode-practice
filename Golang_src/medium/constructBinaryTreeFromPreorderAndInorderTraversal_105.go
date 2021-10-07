package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 89.52%
// Memory Usage: 4.3 MB, less than 69.52%
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	idx := findRoot(inorder, preorder[0])
	root.Left = buildTree(preorder[1:len(inorder[:idx])+1], inorder[:idx])
	root.Right = buildTree(preorder[len(inorder[:idx])+1:], inorder[idx+1:])

	return root
}

// pre: D -> L -> R
//  in: L -> D -> R

func findRoot(inorder []int, target int) int {
	for i := range inorder {
		if inorder[i] == target {
			return i
		}
	}

	return -1
}
