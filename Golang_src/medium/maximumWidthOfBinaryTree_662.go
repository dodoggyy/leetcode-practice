package medium

import "math"

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 13 ms, faster than 16.28%
// Memory Usage: 4.4 MB, less than 100.00%
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	root.Val = 1
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		maxVal, minVal := math.MinInt64, math.MaxInt64
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			minVal = min(minVal, node.Val)
			maxVal = max(maxVal, node.Val)

			if node.Left != nil {
				node.Left.Val = node.Val * 2
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*2 + 1
				queue = append(queue, node.Right)
			}
		}

		result = max(result, maxVal-minVal+1)
	}

	return result
}
