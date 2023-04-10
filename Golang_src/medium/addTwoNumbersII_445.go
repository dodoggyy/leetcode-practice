package medium

// Use stack:
// Time Complexity: O(max(m,n))
// Space Complexity:O(m+n)
// Runtime: 2 ms, faster than 97.73%
// Memory Usage: 5.4 MB, less than 44.32%
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{}, []int{}
	cur1, cur2 := l1, l2
	for cur1 != nil {
		stack1 = append(stack1, cur1.Val)
		cur1 = cur1.Next
	}
	for cur2 != nil {
		stack2 = append(stack2, cur2.Val)
		cur2 = cur2.Next
	}

	carry := 0
	dummy := &ListNode{Val: -1}
	for len(stack1) > 0 || len(stack2) > 0 || carry != 0 {
		tmp1, tmp2 := 0, 0
		if len(stack1) > 0 {
			tmp1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			tmp2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		node := &ListNode{Val: (tmp1 + tmp2 + carry) % 10}
		carry = (tmp1 + tmp2 + carry) / 10
		node.Next = dummy.Next
		dummy.Next = node
	}

	return dummy.Next
}
