package easy

// Use reverse linkedlist & two pointer:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 160 ms, faster than 95.94%
// Memory Usage: 10.9 MB, less than 54.26%
func isPalindromeLinkedList(head *ListNode) bool {
	slow, fast := head, head

	// O(n)
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// O(n)
	half := revertLinkedList(slow)

	// O(n)
	cur := head
	for half != nil {
		if cur.Val != half.Val {
			return false
		}
		cur = cur.Next
		half = half.Next
	}

	return true
}

func revertLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
