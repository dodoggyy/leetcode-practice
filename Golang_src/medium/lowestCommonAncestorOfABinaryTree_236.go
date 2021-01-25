package medium

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 98.17%
// Memory Usage: 7.1 MB, less than 94.82%
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 900 ms, faster than 5.18%
// Memory Usage: 6.7 MB, less than 98.78%
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	queue := []*TreeNode{root}

	result := root
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if containsNode(cur, p) && containsNode(cur, q) {
			result = cur
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return result
}

func containsNode(root, target *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == target {
		return true
	}
	return containsNode(root.Left, target) || containsNode(root.Right, target)
}
