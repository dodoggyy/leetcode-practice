package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use divide and conquer:
// Time Complexity: O(nlogk)
// Space Complexity: O(1)
// Runtime: 8 ms, faster than 95.55%
// Memory Usage: 6.1 MB, less than 33.02%
func mergeKLists(lists []*ListNode) *ListNode {
	var mergeTwoList func(*ListNode, *ListNode) *ListNode
	mergeTwoList = func(a, b *ListNode) *ListNode {
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		for a != nil && b != nil {
			if a.Val < b.Val {
				a.Next = mergeTwoList(a.Next, b)
				return a
			}
			b.Next = mergeTwoList(a, b.Next)
			return b
		}
		return nil
	}

	var merge func([]*ListNode, int, int) *ListNode
	merge = func(lists []*ListNode, l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		if l > r {
			return nil
		}
		mid := (l + r) >> 1
		return mergeTwoList(merge(lists, l, mid), merge(lists, mid+1, r))
	}

	return merge(lists, 0, len(lists)-1)
}
