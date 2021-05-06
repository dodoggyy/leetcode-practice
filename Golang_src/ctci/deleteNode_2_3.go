package ctci

// Use shift data:
// Time Complexity: O(1)
// Space Complexity:O(1)
func deleteNode(node *ListNode) {
	if node == nil || node.next == nil {
		return
	}

	next := node.next
	node.val = next.val
	node.next = node.next.next
}
