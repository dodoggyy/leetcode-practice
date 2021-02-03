package medium

// Use DFS:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 16 ms, faster than 42.72%
// Memory Usage: 4.4 MB, less than 5.83%
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	count := pathStartSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)

	return count
}

func pathStartSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := 0

	if root.Val == sum {
		result++
	}

	return result + pathStartSum(root.Left, sum-root.Val) + pathStartSum(root.Right, sum-root.Val)
}

// Use prefix sum:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 4 ms, faster than 98.06%
// Memory Usage: 5 MB, less than 5.83%
func pathSum2(root *TreeNode, sum int) int {
	hashMap := map[int]int{0: 1}

	return helper(root, 0, sum, hashMap)
}

func helper(node *TreeNode, curSum, sum int, hashMap map[int]int) int {
	if node == nil {
		return 0
	}
	curSum += node.Val
	result := hashMap[curSum-sum]

	hashMap[curSum]++

	result += helper(node.Left, curSum, sum, hashMap)
	result += helper(node.Right, curSum, sum, hashMap)

	hashMap[curSum]--

	return result
}
