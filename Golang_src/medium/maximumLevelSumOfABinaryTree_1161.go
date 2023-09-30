package medium

import "math"

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 99 ms, faster than 78.65%
// Memory Usage: 7.71 MB, less than 76.40%
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := math.MinInt32
	result := 0
	level := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		tmp := 0
		level++
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if tmp > maxSum {
			maxSum = tmp
			result = level

		}
	}

	return result
}
