package medium

// Use simulation:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.36 MB, less than 99.19%
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}
	newList := &ListNode{}
	pre, cur := dummy, head
	tmp := newList
	for cur != nil {
		if cur.Val < x {
			tmp.Next = cur
			tmp = tmp.Next
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}

	tmp.Next = dummy.Next

	return newList.Next
}
