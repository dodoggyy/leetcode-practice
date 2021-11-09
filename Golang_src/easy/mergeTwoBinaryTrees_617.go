package easy

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 32 ms, faster than 30.72%
// Memory Usage: 8.6 MB, less than 92.81%
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var result *TreeNode

	if t1 == nil && t2 == nil {
		return nil
	}

	tmp1, tmp2 := 0, 0
	isNil1, isNil2 := true, true
	if t1 != nil {
		tmp1 = t1.Val
		isNil1 = false
	}
	if t2 != nil {
		tmp2 = t2.Val
		isNil2 = false
	}
	node := &TreeNode{Val: tmp1 + tmp2}
	result = node

	if isNil1 {
		result.Left = mergeTrees(nil, t2.Left)
		result.Right = mergeTrees(nil, t2.Right)
	} else if isNil2 {
		result.Left = mergeTrees(t1.Left, nil)
		result.Right = mergeTrees(t1.Right, nil)
	} else {
		result.Left = mergeTrees(t1.Left, t2.Left)
		result.Right = mergeTrees(t1.Right, t2.Right)
	}

	return result
}

// Optimize DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 28 ms, faster than 90.20%
// Memory Usage: 8.8 MB, less than 92.81%
func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees2(t1.Left, t2.Left)
	t1.Right = mergeTrees2(t1.Right, t2.Right)

	return t1
}
