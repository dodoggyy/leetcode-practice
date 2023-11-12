package medium

import "sort"

// Use array sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
// Runtime: 52 ms, faster than 65.33%
// Memory Usage: 8.13 MB, less than 25.00%
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	sort.Ints(list)
	result := &ListNode{Val: list[0]}
	cur := result
	for i := 1; i < len(list); i++ {
		cur.Next = &ListNode{Val: list[i]}
		cur = cur.Next
	}

	return result
}

// Use divide and conquer (bottom up):
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 55 ms, faster than 50.00%
// Memory Usage: 7.13 MB, less than 95.00%
func sortList2(head *ListNode) *ListNode {
	merge := func(head1, head2 *ListNode) *ListNode {
		dummy := &ListNode{}
		cur := dummy
		tmp1, tmp2 := head1, head2
		for tmp1 != nil && tmp2 != nil {
			if tmp1.Val <= tmp2.Val {
				cur.Next = tmp1
				tmp1 = tmp1.Next
			} else {
				cur.Next = tmp2
				tmp2 = tmp2.Next
			}
			cur = cur.Next
		}
		if tmp1 != nil {
			cur.Next = tmp1
		} else {
			cur.Next = tmp2
		}
		return dummy.Next
	}

	if head == nil {
		return head
	}
	size := 0
	for cur := head; cur != nil; cur = cur.Next {
		size++
	}

	dummy := &ListNode{Next: head}
	for subSize := 1; subSize < size; subSize <<= 1 {
		pre, cur := dummy, dummy.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subSize && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subSize && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			pre.Next = merge(head1, head2)

			for pre.Next != nil {
				pre = pre.Next
			}
			cur = next
		}
	}

	return dummy.Next
}
