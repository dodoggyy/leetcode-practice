package medium

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 225 ms, faster than 57.29%
// Memory Usage: 9.5 MB, less than 81.25%
func swapNodes(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	var node *ListNode
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}

	node = fast

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	node.Val, slow.Val = slow.Val, node.Val

	return head
}
