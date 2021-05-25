package easy

// Use two pointers:
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// Runtime: 40 ms, faster than 84.90%
// Memory Usage: 8.6 MB, less than 26.30%
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	return curA
}
