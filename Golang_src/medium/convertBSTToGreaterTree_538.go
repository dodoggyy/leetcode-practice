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
// Runtime: 224 ms, faster than 25.49%
// Memory Usage: 248.9 MB, less than 50.98%
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	traversal(root, &sum)

	return root
}

func traversal(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	traversal(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	traversal(root.Left, sum)
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
