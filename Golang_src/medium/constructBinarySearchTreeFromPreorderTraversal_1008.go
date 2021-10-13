package medium

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.4 MB, less than 33.33%
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	findR := func(target int) int {
		for i := range preorder {
			if preorder[i] > target {
				return i
			}
		}
		return -1
	}

	root := &TreeNode{Val: preorder[0]}
	idxR := findR(preorder[0])

	if idxR > 0 {
		root.Left = bstFromPreorder(preorder[1:idxR])
		root.Right = bstFromPreorder(preorder[idxR:])
	} else {
		root.Left = bstFromPreorder(preorder[1:])
	}

	return root
}

// Use DFS 2:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 11.11%
// Memory Usage: 2.5 MB, less than 11.11%
func bstFromPreorder2(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	findR := func(target int) int {
		for i := range preorder {
			if preorder[i] > target {
				return i
			}
		}
		return len(preorder)
	}

	idxR := findR(preorder[0])

	root.Left = bstFromPreorder(preorder[1:idxR])
	root.Right = bstFromPreorder(preorder[idxR:])

	return root
}
