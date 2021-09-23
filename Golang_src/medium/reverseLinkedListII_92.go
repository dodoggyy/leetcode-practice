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
