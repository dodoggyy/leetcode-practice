package easy

// Use inorder traversal:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 30.49%
// Memory Usage: 2.2 MB, less than 100.00%
func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	val := []int{}
	cur := dummy

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		val = append(val, node.Val)
		dfs(node.Right)
	}

	dfs(root)

	for _, v := range val {
		cur.Right = &TreeNode{Val: v}
		cur = cur.Right
	}

	return dummy.Right
}

// Use Traversal with Relinking:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 1 ms, faster than 32.93%
// Memory Usage: 2.2 MB, less than 100.00%
func increasingBST2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		cur.Right = node
		cur = cur.Right
		node.Left = nil
		dfs(node.Right)
	}

	dfs(root)

	return dummy.Right
}
