package easy

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 97.54%
// Memory Usage: 3.8 MB, less than 73.72%
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
