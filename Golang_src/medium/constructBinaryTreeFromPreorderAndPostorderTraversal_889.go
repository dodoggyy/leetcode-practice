package medium

// Use recursive:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 4 ms, faster than 77.27%
// Memory Usage: 3.3 MB, less than 50.00%
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	idxLeftRoot := findIdxLeftSubTreeRoot(postorder, preorder[1])

	l := constructFromPrePost(preorder[1:idxLeftRoot+2], postorder[:idxLeftRoot+1])
	r := constructFromPrePost(preorder[idxLeftRoot+2:], postorder[idxLeftRoot+1:len(postorder)-1])
	root.Left, root.Right = l, r

	return root
}

//  pre: D -> L -> R
// post: L -> R -> D

// e.g. [1, 2, 3, 4, 5, 6, 7]
//  pre: [1] + ["2", 4, 5] + [3, 6, 7]
// post: [4, 5, "2"] + [6, 7, 3] + [1]

func findIdxLeftSubTreeRoot(postorder []int, target int) int {
	for i := range postorder {
		if postorder[i] == target {
			return i
		}
	}
	return 0
}
