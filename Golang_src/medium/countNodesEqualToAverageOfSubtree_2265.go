package medium

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 6 ms, faster than 54.84%
// Memory Usage: 4.41 MB, less than 85.48%
func averageOfSubtree(root *TreeNode) int {
	result := 0

	var helper func(node *TreeNode) (int, int)
	helper = func(node *TreeNode) (int, int) {
		sum, cnt := node.Val, 1
		if node.Left != nil {
			sumSub, cntSub := helper(node.Left)
			sum += sumSub
			cnt += cntSub
		}
		if node.Right != nil {
			sumSub, cntSub := helper(node.Right)
			sum += sumSub
			cnt += cntSub
		}
		if sum/cnt == node.Val {
			result++
		}
		return sum, cnt
	}

	helper(root)

	return result
}
