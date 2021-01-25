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
// Runtime: 12 ms, faster than 100.00%
// Memory Usage: 6.8 MB, less than 5.29%
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 20 ms, faster than 81.18%
// Memory Usage: 6.7 MB, less than 93.53%
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	cur := root
	pVal, qVal := p.Val, q.Val

	for cur != nil {
		curVal := cur.Val

		if curVal > pVal && curVal > qVal {
			cur = cur.Left
		} else if curVal < pVal && curVal < qVal {
			cur = cur.Right
		} else {
			return cur
		}
	}

	return nil
}
