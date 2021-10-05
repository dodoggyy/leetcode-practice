package medium

// Use level order traversal:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 48 ms, faster than 46.22%
// Memory Usage: 2.3 MB, less than 48.28%
func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		tmp := []int{}
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, tmp[len(tmp)-1])
	}

	return result
}
