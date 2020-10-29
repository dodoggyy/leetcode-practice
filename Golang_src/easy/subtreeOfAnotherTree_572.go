package easy

import (
	"strconv"
	"strings"
)

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use node comparison:
// Time Complexity: O(m*n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 99.07%
// Memory Usage: 6.2 MB, less than 10.19%
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	return treeCompare(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func treeCompare(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return treeCompare(p.Left, q.Left) && treeCompare(p.Right, q.Right)
}

// Use Preorder Traversal :
// Time Complexity: O(m*n)
// Space Complexity:O(max(m,n))
// Runtime: 16 ms, faster than 93.98%
// Memory Usage: 7 MB, less than 10.19%
func isSubtree2(s *TreeNode, t *TreeNode) bool {
	strS := preorder(s)
	strT := preorder(t)

	return strings.Contains(strS, strT)
}

func preorder(root *TreeNode) string {
	if root == nil {
		return "*"
	}

	return "#" + strconv.Itoa(root.Val) + preorder(root.Left) + preorder(root.Right)
}
