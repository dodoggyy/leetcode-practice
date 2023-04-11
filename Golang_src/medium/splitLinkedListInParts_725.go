package medium

// Use split linkedlist:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100%
// Memory Usage: 2.8 MB, less than 76.92%
func splitListToParts(head *ListNode, k int) []*ListNode {
	result := make([]*ListNode, k)

	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}

	each := cnt / k
	mod := cnt % k

	cur = head
	for i := 0; cur != nil && i < k; i++ {
		result[i] = cur

		size := each
		if i < mod {
			size++
		}

		for i := 0; i < size-1; i++ {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}

	return result
}
