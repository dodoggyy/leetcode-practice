package medium

// Use dfs:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 9 ms, faster than 14.92%
// Memory Usage: 4.7 MB, less than 59.27%
func pathSumII(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	keep := []int{}

	var dfs func(node *TreeNode, remain int)
	dfs = func(node *TreeNode, remain int) {
		if node == nil {
			return
		}
		keep = append(keep, node.Val)
		if node.Left == nil && node.Right == nil && remain == node.Val {
			tmp := make([]int, len(keep))
			copy(tmp, keep)
			result = append(result, tmp)
		}
		dfs(node.Left, remain-node.Val)
		dfs(node.Right, remain-node.Val)
		keep = keep[:len(keep)-1]
	}

	dfs(root, targetSum)

	return result
}
