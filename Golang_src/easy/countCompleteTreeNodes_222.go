package easy

// Use binary search:
// Time Complexity: O(logn*logn)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 77.78%
// Memory Usage: 7.17 MB, less than 51.47%
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levelL, levelR := 0, 0
	nodeL, nodeR := root, root
	for nodeL != nil {
		levelL++
		nodeL = nodeL.Left
	}
	for nodeR != nil {
		levelR++
		nodeR = nodeR.Right
	}
	if levelL == levelR {
		return 1<<levelL - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
