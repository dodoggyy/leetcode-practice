package easy

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 240 ms, faster than 48.33%
// Memory Usage: 18.3 MB, less than 10.00%
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	}

	return min(left, right) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 248 ms, faster than 30.83%
// Memory Usage: 21.7 MB, less than 10.00%
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	result := 0
	level := 1

	for len(queue) > 0 {
		result++
		level = len(queue)

		for level > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return result
			}
			level--
		}
	}

	return result
}
