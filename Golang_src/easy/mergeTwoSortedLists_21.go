package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use recursive:
// Time Complexity: O(max(m+n))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.6 MB, less than 49.24%
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

	return nil
}

// Use iterative:
// Time Complexity: O(max(m+n))
// Space Complexity:O(m+n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.5 MB, less than 49.24%
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	cur := &ListNode{Val: -1, Next: nil}
	head := cur

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return head.Next
}
