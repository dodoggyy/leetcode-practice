package easy

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 8 ms, faster than 90.24%
// Memory Usage: 6.2 MB, less than 13.41%
func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var sum float64 = 0

		for _, node := range queue {
			sum += float64(node.Val)
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, sum/float64(size))
	}

	return result
}
