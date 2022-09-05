package medium

type Node2 struct {
	Val      int
	Children []*Node2
}

// Use BFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 5 ms, faster than 49.18%
// Memory Usage: 4.4 MB, less than 70.49%
func levelOrderTraversal(root *Node2) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*Node2{root}
	for len(queue) > 0 {
		tmp := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			tmp = append(tmp, node.Val)

			for _, v := range node.Children {
				queue = append(queue, v)
			}

		}

		result = append(result, tmp)
	}
	return result
}
