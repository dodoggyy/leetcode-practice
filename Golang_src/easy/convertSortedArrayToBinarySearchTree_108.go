package easy

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use recursive:
// Time Complexity: O(n)
// Space Complexity: O(n)
// Runtime: 96 ms, faster than 91.25%
// Memory Usage: 233.2 MB, less than 31.31%
func sortedArrayToBST(nums []int) *TreeNode {
	return toBST(nums, 0, len(nums)-1)
}

func toBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = toBST(nums, start, mid-1)
	node.Right = toBST(nums, mid+1, end)

	return node
}
