package medium

type Node3 struct {
	Val   int
	Left  *Node3
	Right *Node3
	Next  *Node3
}

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 3 ms, faster than 67.48%
// Memory Usage: 6.43 MB, less than 58.25%
func connect(root *Node3) *Node3 {
	if root == nil {
		return root
	}

	queue := []*Node3{root}
	for len(queue) > 0 {
		size := len(queue)
		var pre *Node3
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if pre != nil {
				pre.Next = node
			}
			pre = node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
