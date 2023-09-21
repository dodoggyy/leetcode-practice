package medium

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 302 ms, faster than 37.57%
// Memory Usage: 9.18 MB, less than 76.17%
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var pre *ListNode

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	pre.Next = slow.Next

	return head
}
