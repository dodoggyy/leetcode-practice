package ctci

// Use calculate list length:
// Time Complexity: O(n)
// Space Complexity:O(1)
func KFromTail(node *ListNode, k int) *ListNode {
	head := node
	tmp := node
	cnt := 1
	for tmp != nil {
		cnt++
		tmp = tmp.next
	}

	for i := 0; i < cnt-k-1; i++ {
		head = head.next
	}

	return head
}

// Use two pointer:
// Time Complexity: O(n)
// Space Complexity:O(1)
func KFromTail2(node *ListNode, k int) *ListNode {
	slow, fast := node, node
	for i := 0; i <= k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.next
	}
	for fast != nil {
		slow = slow.next
		fast = fast.next
	}

	return slow
}
