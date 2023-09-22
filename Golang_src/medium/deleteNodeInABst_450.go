package medium

// Use recursive :
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 11 ms, faster than 68.94%
// Memory Usage: 6.79 MB, less than 93.19%
func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		if root.Left == nil {
			return root.Right
		}
		return root.Left
	default:
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
