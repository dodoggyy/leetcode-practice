package medium

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 4 ms, faster than 100.00%
// Memory Usage: 5.7 MB, less than 70.59%
func findBottomLeftValue(root *TreeNode) int {
	var node *TreeNode
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node = queue[0]
			queue = queue[1:]

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}

	return node.Val
}
