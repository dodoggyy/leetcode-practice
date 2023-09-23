package medium

// Use reverse + iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 220 ms, faster than 21.65%
// Memory Usage: 9.79 MB, less than 70.47%
func pairSum(head *ListNode) int {
	result := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	deepCopy := func(node *ListNode) *ListNode {
		pre := &ListNode{}
		tmp := node
		tmpCopy := pre
		for tmp != nil {
			tmpCopy.Next = &ListNode{
				Val: tmp.Val,
			}
			tmp = tmp.Next
			tmpCopy = tmpCopy.Next
		}

		return pre.Next
	}

	reverse := func(node *ListNode) *ListNode {
		var pre *ListNode
		for node != nil {
			tmp := node.Next
			node.Next = pre
			pre = node
			node = tmp
		}

		return pre
	}

	headCopy := deepCopy(head)
	headCopy = reverse(headCopy)

	for head != nil && headCopy != nil {
		result = max(result, head.Val+headCopy.Val)
		head = head.Next
		headCopy = headCopy.Next
	}

	return result
}

// Use optimize reverse + iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 186 ms, faster than 65.35%
// Memory Usage: 8.93 MB, less than 70.47%
func pairSum2(head *ListNode) int {
	result := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	reverse := func(node *ListNode) *ListNode {
		var pre *ListNode
		for node != nil {
			tmp := node.Next
			node.Next = pre
			pre = node
			node = tmp
		}

		return pre
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	last := slow.Next
	slow.Next = nil
	last = reverse(last)

	for head != nil && last != nil {
		result = max(result, head.Val+last.Val)
		head = head.Next
		last = last.Next
	}

	return result
}
