package medium

import "math"

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 32 ms, faster than 8.90%
// Memory Usage: 5.6 MB, less than 28.23%
func isValidBST(root *TreeNode) bool {
	return helpers(root, math.MinInt64, math.MaxInt64)
}

func helpers(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helpers(root.Right, root.Val, upper) && helpers(root.Left, lower, root.Val)
}

// Use inorder traversal:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 44.37%
// Memory Usage: 5.8 MB, less than 12.34%
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	cur := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= cur {
			return false
		}
		cur = root.Val
		root = root.Right
	}
	return true
}
