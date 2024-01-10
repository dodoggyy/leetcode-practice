package medium

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 167 ms, faster than 77.78%
// Memory Usage: 42.60 MB, less than 33.33%
func amountOfTime(root *TreeNode, start int) int {
	maxDist := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		depth := 0
		if node == nil {
			return depth
		}

		depthL := traverse(node.Left)
		depthR := traverse(node.Right)

		if node.Val == start {
			maxDist = max(depthL, depthR)
			depth = -1
		} else if depthL >= 0 && depthR >= 0 {
			depth = max(depthL, depthR) + 1
		} else {
			dist := abs(depthL) + abs(depthR)
			maxDist = max(maxDist, dist)
			depth = min(depthL, depthR) - 1
		}

		return depth
	}

	traverse(root)
	return maxDist
}
