package medium

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 397 ms, faster than 12.58%
// Memory Usage: 15.9 MB, less than 69.18%
func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	sum := 0
	for head != nil {
		if head.Val == 0 {
			if sum != 0 {
				node := &ListNode{Val: sum}
				sum = 0
				cur.Next = node
				cur = cur.Next
			}
		} else {
			sum += head.Val
		}
		head = head.Next
	}

	return dummy.Next
}
