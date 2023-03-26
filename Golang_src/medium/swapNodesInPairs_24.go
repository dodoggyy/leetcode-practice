package medium

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 51.87%
// Memory Usage: 2.1 MB, less than 56.26%
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		l1, l2 := cur.Next, cur.Next.Next
		tmp := l2.Next

		l1.Next, l2.Next = tmp, l1
		cur.Next = l2
		cur = l1
	}

	return dummy.Next
}

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 51.87%
// Memory Usage: 2.1 MB, less than 56.26%
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headNew := head.Next
	head.Next = swapPairs2(headNew.Next)
	headNew.Next = head

	return headNew
}
