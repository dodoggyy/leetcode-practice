package medium

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 6.79%
// Memory Usage: 3.3 MB, less than 17.65%
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	evenHead := head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}
