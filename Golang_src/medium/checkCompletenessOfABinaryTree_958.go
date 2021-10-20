package medium

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 3 ms, faster than 50.00%
// Memory Usage: 3.4 MB, less than 9.26%
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}

	isEnd := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			isEnd = true
		} else {
			if isEnd == true {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return true
}
