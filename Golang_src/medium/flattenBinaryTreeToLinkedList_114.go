package medium

// Use preorder traversal:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.3 MB, less than 6.67%
func flatten(root *TreeNode) {
	list := preorderTraversal3(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left = nil
		pre.Right = cur
	}
}

func preorderTraversal3(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	preorder := []*TreeNode{}
	preorder = append(preorder, root)
	preorder = append(preorder, preorderTraversal3(root.Left)...)
	preorder = append(preorder, preorderTraversal3(root.Right)...)

	return preorder
}

// Use find predecessor node:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.9 MB, less than 49.52%
func flatten2(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left = nil
			cur.Right = next
		}
		cur = cur.Right
	}
}
