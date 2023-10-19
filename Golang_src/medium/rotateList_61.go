package medium

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.61 MB, less than 9.93%
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	size := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		size++
	}
	if size == 1 {
		return head
	}

	k = k % size
	if k == 0 {
		return head
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	pre := slow.Next
	slow.Next = nil
	fast.Next = head

	return pre
}

// Use iterative 2:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.62 MB, less than 9.93%
func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	size := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		size++
	}

	rotate := size - k%size
	if k == size {
		return head
	}
	cur.Next = head
	for rotate > 0 {
		cur = cur.Next
		rotate--
	}
	newHead := cur.Next
	cur.Next = nil

	return newHead
}
