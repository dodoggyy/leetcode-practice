package easy

// Use preorder:
// Time Complexity: O(t1+t2)
// Space Complexity:O(t1+t2)
// Runtime: 6 ms, faster than 6.25%
// Memory Usage: 2.7 MB, less than 21.88%
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var preorder func(root *TreeNode) []int
	preorder = func(root *TreeNode) []int {
		tmp := []int{}
		if root == nil {
			return tmp
		}

		tmp = append(tmp, preorder(root.Left)...)
		if root.Left == nil && root.Right == nil {
			tmp = append(tmp, root.Val)
		}
		tmp = append(tmp, preorder(root.Right)...)

		return tmp
	}

	tmp1, tmp2 := preorder(root1), preorder(root2)
	if len(tmp1) != len(tmp2) {
		return false
	}

	for i := range tmp1 {
		if tmp1[i] != tmp2[i] {
			return false
		}
	}

	return true
}
