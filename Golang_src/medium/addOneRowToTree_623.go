package medium

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 93.88%%
// Memory Usage: 5.6 MB, less than 83.67%
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	if depth == 2 {
		root.Left = &TreeNode{
			Val:  val,
			Left: root.Left,
		}
		root.Right = &TreeNode{
			Val:   val,
			Right: root.Right,
		}
	} else {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
	}

	return root
}
