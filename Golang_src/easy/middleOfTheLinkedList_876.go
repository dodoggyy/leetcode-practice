package easy

// Use 2 pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}

	return slow
}
