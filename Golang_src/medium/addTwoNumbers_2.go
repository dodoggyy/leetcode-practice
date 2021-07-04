package medium

// Use iterative:
// Time Complexity: O(max(m,n))
// Space Complexity:O(max(m,n))
// Runtime: 12 ms, faster than 59.26%
// Memory Usage: 4.8 MB, less than 93.91%
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result

	remain := 0

	// O(max(m,n))
	for l1 != nil || l2 != nil || remain != 0 {
		if l1 != nil {
			remain += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remain += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: remain % 10}
		cur = cur.Next
		remain /= 10
	}

	return result.Next
}
