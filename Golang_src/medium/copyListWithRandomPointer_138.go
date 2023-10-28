package medium

type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

// Use hashmap:
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

// Use hashmap 2:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 70.25%
// Memory Usage: 3.72 MB, less than 5.83%
func copyRandomList2(head *RNode) *RNode {
	if head == nil {
		return head
	}
	cur := head
	hashmap := map[*RNode]*RNode{}
	for cur != nil {
		hashmap[cur] = &RNode{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		node := hashmap[cur]
		node.Next = hashmap[cur.Next]
		node.Random = hashmap[cur.Random]
		cur = cur.Next
	}

	return hashmap[head]
}
