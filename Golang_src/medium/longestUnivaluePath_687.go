package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 68 ms, faster than 100.00%
// Memory Usage: 6.8 MB, less than 100.00%
func longestUnivaluePath(root *TreeNode) int {
	result := 0
	subPath(root, &result)

	return result
}

func subPath(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	resultL, resultR := subPath(root.Left, result), subPath(root.Right, result)

	pathL, pathR := 0, 0

	if root.Left != nil && root.Left.Val == root.Val {
		pathL = resultL + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		pathR = resultR + 1
	}

	*result = max(*result, pathL+pathR)

	return max(pathL, pathR)
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 108 ms, faster than 6.67%
// Memory Usage: 7.3 MB, less than 100.00%
func longestUnivaluePath2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = max(result, getLongestPathByValue(node.Left, node.Val)+getLongestPathByValue(node.Right, node.Val))

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func getLongestPathByValue(root *TreeNode, value int) int {
	maxDepth := 0
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := []int{1}

	for len(queue) > 0 && len(depth) > 0 {
		node := queue[0]
		queue = queue[1:]
		pollDepth := depth[0]
		depth = depth[1:]

		if node.Val == value {
			maxDepth = max(maxDepth, pollDepth)
			if node.Left != nil {
				queue = append(queue, node.Left)
				depth = append(depth, pollDepth+1)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				depth = append(depth, pollDepth+1)
			}
		}
	}

	return maxDepth
}
