package medium

// Use backtracking:
// Time Complexity: O(n*Gn)
// Space Complexity:O(n*Gn)
// Gn: Catalan number of n
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 0.0 MB, less than 100%
func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return nil
	}

	return helperTreeGenerate(1, n)
}

func helperTreeGenerate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil} // need set as nil
	}
	result := []*TreeNode{}

	for i := start; i <= end; i++ {
		leftTree := helperTreeGenerate(start, i-1)
		rightTree := helperTreeGenerate(i+1, end)

		for _, left := range leftTree {
			for _, right := range rightTree {
				curTree := &TreeNode{Val: i}
				curTree.Left = left
				curTree.Right = right
				result = append(result, curTree)
			}
		}
	}

	return result
}
