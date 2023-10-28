package medium

// Use Iterative Link Reversal:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 37.99%
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	pre := dummy

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next

	for i := 0; i < right-left; i++ {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}

	return dummy.Next
}

// Use Iterative Link Reversal:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 79.30%
// Memory Usage: 2.14 MB, less than 11.62%
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	reverse := func(head *ListNode) *ListNode {
		var pre *ListNode

		for head != nil {
			tmp := head.Next
			head.Next = pre
			pre = head
			head = tmp
		}
		return pre
	}

	dummy := &ListNode{Next: head}
	pre := dummy

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	nodeRight := pre
	for i := left; i <= right; i++ {
		nodeRight = nodeRight.Next
	}

	nodeLeft := pre.Next
	cur := nodeRight.Next

	pre.Next = nil
	nodeRight.Next = nil

	nodeRight = reverse(nodeLeft)

	pre.Next = nodeRight
	nodeLeft.Next = cur

	return dummy.Next
}
