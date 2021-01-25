package easy

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.2 MB, less than 100.00%
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)

	return root
}

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.2 MB, less than 100.00%
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}

	return root
}
