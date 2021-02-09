package medium

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.8 MB, less than 100.00%
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
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
			if node != nil {
				tmp = append(tmp, node.Val)
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

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3 MB, less than 34.47%
func levelOrder2(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(result) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	return result
}
