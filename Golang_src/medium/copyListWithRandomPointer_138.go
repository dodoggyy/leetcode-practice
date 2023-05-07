package medium

type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.5 MB, less than 54.5%
func copyRandomList(head *RNode) *RNode {
	newHead := &RNode{}

	hashmap := map[*RNode]*RNode{}

	cur := head
	curNew := newHead
	for cur != nil {
		node := &RNode{Val: cur.Val}

		hashmap[cur] = node

		curNew.Next = node
		curNew = curNew.Next
		cur = cur.Next
	}

	curNew = newHead.Next
	cur = head
	for cur != nil {
		curNew.Random = hashmap[cur.Random]
		cur = cur.Next
		curNew = curNew.Next
	}

	return newHead.Next
}
