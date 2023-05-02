package medium

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
	dfsRob(root)

	return max(hashmapRob[root], hashmapNotRob[root])
}

func dfsRob(root *TreeNode) {
	if root == nil {
		return
	}
	dfsRob(root.Left)
	dfsRob(root.Right)

	hashmapRob[root] = root.Val + hashmapNotRob[root.Left] + hashmapNotRob[root.Right]

	hashmapNotRob[root] = max(hashmapRob[root.Left], hashmapNotRob[root.Left]) + max(hashmapRob[root.Right], hashmapNotRob[root.Right])
}

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 6 ms, faster than 58.33%
// Memory Usage: 5.5 MB, less than 51.4%
func rob6(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var helper func(cur *TreeNode) []int
	helper = func(cur *TreeNode) []int {
		if cur == nil {
			return []int{0, 0}
		}
		l, r := helper(cur.Left), helper(cur.Right)

		// 0: not rob, 1: rob
		rob := cur.Val + l[0] + r[0]
		noRob := max(l[0], l[1]) + max(r[0], r[1])

		return []int{noRob, rob}
	}

	result := helper(root)

	return max(result[0], result[1])
}
