package medium

import "math"

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 9 ms, faster than 40.63%
// Memory Usage: 6.06 MB, less than 42.19%
func largestValues(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		tmp := math.MinInt32
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > tmp {
				tmp = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, tmp)
	}

	return result
}
