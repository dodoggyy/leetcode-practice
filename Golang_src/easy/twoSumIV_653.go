package easy

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use BST:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 84 ms, faster than 5.75%
// Memory Usage: 8.9 MB, less than 9.20%
func findTarget(root *TreeNode, k int) bool {
	nums := inorderTraversal(root)

	start, end := 0, len(nums)-1

	for start < end {
		sum := nums[start] + nums[end]
		if sum > k {
			end--
		} else if sum < k {
			start++
		} else {
			return true
		}
	}

	return false
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	inorder := []int{}
	inorder = append(inorder, inorderTraversal(root.Left)...)
	inorder = append(inorder, root.Val)
	inorder = append(inorder, inorderTraversal(root.Right)...)

	return inorder
}

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 44 ms, faster than 12.64%
// Memory Usage: 6.7 MB, less than 86.21%
func findTarget2(root *TreeNode, k int) bool {
	hashmap := map[int]bool{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if _, exist := hashmap[k-node.Val]; exist {
			return true
		}
		hashmap[node.Val] = true
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}
