package hard

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(h)
// Runtime: 4 ms, faster than 93.06%
// Memory Usage: 3.7 MB, less than 100.00%
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		tmp := tail.Next
		head, tail = reverseSpecific(head, tail)
		pre.Next = head
		tail.Next = tmp
		pre = tail
		head = tail.Next
	}

	return dummy.Next
}

func reverseSpecific(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head

	for pre != tail { // reverse until tail rather cur is nil
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	// after reverse: head: tail, tail: head
	return tail, head
}
