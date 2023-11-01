package easy

// Use inorder:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 51.06%
// Memory Usage: 7.6 MB, less than 6.38%
func findMode(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	order := inorder(root)

	if len(order) == 1 {
		return []int{order[0]}
	}

	result := []int{}

	cur := 1
	max := 1

	for i := 1; i < len(order); i++ {
		if order[i] == order[i-1] {
			cur++
		} else {
			if cur == max {
				result = append(result, order[i-1])
			} else if cur > max {
				result = []int{order[i-1]}
				max = cur
			}
			cur = 1
		}
	}

	if cur == max {
		result = append(result, order[len(order)-1])
	} else if cur > max {
		result = []int{order[len(order)-1]}
	}

	return result
}

func inorder3(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	order := []int{}
	order = append(order, inorder3(root.Left)...)
	order = append(order, root.Val)
	order = append(order, inorder3(root.Right)...)

	return order
}

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 51.06%
// Memory Usage: 6 MB, less than 53.19%
func findMode2(root *TreeNode) []int {
	maxCount := 0
	hashmap := map[int]int{}

	result := []int{}

	inorder2(root, hashmap, &maxCount)

	for k, v := range hashmap {
		if v == maxCount {
			result = append(result, k)
		}
	}

	return result
}

func inorder2(root *TreeNode, hashmap map[int]int, maxCount *int) {

	if root == nil {
		return
	}

	hashmap[root.Val]++
	*maxCount = max(*maxCount, hashmap[root.Val])

	inorder2(root.Left, hashmap, maxCount)
	inorder2(root.Right, hashmap, maxCount)

}

// Use dfs:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 86.36%
// Memory Usage: 6.02 MB, less than 100.00%
func findMode3(root *TreeNode) []int {
	base, cnt, maxCnt := 0, 0, 0

	result := []int{}

	update := func(x int) {
		if x == base {
			cnt++
		} else {
			base = x
			cnt = 1
		}

		if cnt == maxCnt {
			result = append(result, base)
		} else if cnt > maxCnt {
			maxCnt = cnt
			result = []int{base}
		}
	}

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		update(node.Val)
		inorder(node.Right)
	}

	inorder(root)

	return result
}
