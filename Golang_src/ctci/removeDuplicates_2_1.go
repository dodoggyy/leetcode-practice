package ctci

type ListNode struct {
	val  int
	pre  *ListNode
	next *ListNode
}

// Use hashset:
// Time Complexity: O(N)
// Space Complexity:O(N)
func removeDuplicate(node *ListNode) {
	hashset := map[int]bool{}
	cur := node
	var previous *ListNode
	for cur != nil {
		if _, ok := hashset[cur.val]; ok {
			previous.next = cur.next
		} else {
			hashset[cur.val] = true
			previous = cur
		}
		cur = cur.next
	}
}
