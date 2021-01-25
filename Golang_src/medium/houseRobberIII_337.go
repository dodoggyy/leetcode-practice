package medium

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var hashmap map[*TreeNode]int = make(map[*TreeNode]int)

var hashmapRob map[*TreeNode]int = make(map[*TreeNode]int)
var hashmapNotRob map[*TreeNode]int = make(map[*TreeNode]int)

// Use recursive without memoraization:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 676 ms, faster than 15.60%
// Memory Usage: 5.2 MB, less than 100.00%
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val1 := root.Val
	if root.Left != nil {
		val1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob(root.Right.Left) + rob(root.Right.Right)
	}
	val2 := rob(root.Left) + rob(root.Right)

	return max(val1, val2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Use recursive with memoraization:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 60.55%
// Memory Usage: 6.1 MB, less than 15.60%
func rob2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if _, exist := hashmap[root]; exist {
		return hashmap[root]
	}

	val1 := root.Val
	if root.Left != nil {
		val1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob(root.Right.Left) + rob(root.Right.Right)
	}
	val2 := rob(root.Left) + rob(root.Right)

	hashmap[root] = max(val1, val2)

	return hashmap[root]
}

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 60.55%
// Memory Usage: 6.1 MB, less than 15.60%
func rob3(root *TreeNode) int {
	dfs(root)

	return max(hashmapRob[root], hashmapNotRob[root])
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)

	hashmapRob[root] = root.Val + hashmapNotRob[root.Left] + hashmapNotRob[root.Right]

	hashmapNotRob[root] = max(hashmapRob[root.Left], hashmapNotRob[root.Left]) + max(hashmapRob[root.Right], hashmapNotRob[root.Right])
}
