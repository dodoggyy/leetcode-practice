package easy

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 0.0 MB, less than 100%
func isSymmetric(root *TreeNode) bool {
	return compare(root, root)
}

func compare(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return compare(p.Left, q.Right) && compare(p.Right, q.Left)
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.1 MB, less than 5.96%
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode
	queue = append(queue, root)
	queue = append(queue, root)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		q := queue[0]
		queue = queue[1:]

		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left)
		queue = append(queue, q.Right)
		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
	}

	return true
}
